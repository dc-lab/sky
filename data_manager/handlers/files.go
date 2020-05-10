package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/gabriel-vasile/mimetype"

	"github.com/dc-lab/sky/data_manager/config"
	errs "github.com/dc-lab/sky/data_manager/errors"
	"github.com/dc-lab/sky/data_manager/modelapi"
	"github.com/dc-lab/sky/data_manager/modeldb"
	"github.com/dc-lab/sky/data_manager/repo"
)

const maxUploadBufferSize = 8 * 1024 * 1024 // 8 MiB

type FilesService struct {
	Repo   *repo.FilesRepo
	Config *config.Config
}

func (s *FilesService) FileCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var file modeldb.File
		var err error

		if fileId := chi.URLParam(r, "fileId"); fileId != "" {
			file, err = s.Repo.Get(fileId)
		}
		if err != nil {
			log.WithError(err).Error("Not found requested file")
			render.Render(w, r, errs.NotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "file", &file)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *FilesService) getFile(r *http.Request) *modeldb.File {
	return r.Context().Value("file").(*modeldb.File)
}

func (s *FilesService) makeUploadUrl(r *http.Request, id string, uploadToken string) string {
	uri := r.URL
	query := make(url.Values)
	query.Set("token", uploadToken)

	uri.Path += fmt.Sprintf("/%s/data", id)
	uri.RawQuery = query.Encode()
	log.WithFields(log.Fields{
		"token": uploadToken,
		"query": uri.RawQuery,
	}).Info("Debug url")
	return r.Host + uri.String()
}

// CreateFile godoc
// @Summary Create new file
// @Description Create new file metadata, actual file binary should be uploaded later
// @Tags files
// @Accept json
// @Produce json
// @Param file body modelapi.FileRequest true "File metadata"
// @Success 200 {object} modelapi.FileResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /files [post]
func (s *FilesService) CreateFile(w http.ResponseWriter, r *http.Request) {
	var request modelapi.FileRequest
	if err := render.Decode(r, &request); err != nil {
		log.WithError(err).Error("Failed to decode request")
		render.Render(w, r, errs.BadRequest(err))
		return
	}

	file, err := s.Repo.Create(*request.File)
	if err != nil {
		log.WithError(err).Error("Failed to insert new file into db")
		render.Render(w, r, errs.Internal(err))
		return
	}

	var response modelapi.FileResponse
	response.File = &file
	response.UploadUrl = s.makeUploadUrl(r, file.Id, file.UploadToken)
	render.Respond(w, r, response)
}

// GetFile godoc
// @Summary Get existing file metadata
// @Tags files
// @Produce json
// @Param id path string true "File id"
// @Success 200 {object} modelapi.FileResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /files/{id} [get]
func (s *FilesService) GetFile(w http.ResponseWriter, r *http.Request) {
	file := s.getFile(r)
	if file == nil {
		log.Error("Failed to load file")
		render.Render(w, r, errs.NotFound)
		return
	}

	render.Respond(w, r, modelapi.FileResponse{File: file})
}

func fileNotExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// UploadFileData godoc
// @Summary Upload file data
// @Description Upload binary data for existing file metadata. Should not be called directly, use returned upload url instead
// @Tags files
// @Accept mpfd
// @Produce json
// @Param id path string true "File id"
// @Param token query string true "Token returned with metadata"
// @Param file formData file true "File contents"
// @Success 200 {object} modelapi.FileResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 413 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /files/{id}/data [post]
func (s *FilesService) UploadFileData(w http.ResponseWriter, r *http.Request) {
	file := s.getFile(r)
	if file == nil {
		log.Error("Failed to load file")
		render.Render(w, r, errs.NotFound)
		return
	}

	if file.UploadToken != r.URL.Query().Get("token") {
		log.Error("Invalid upload token")
		render.Render(w, r, errs.NotFound)
		return
	}

	if file.Hash != "" {
		log.Error("Could not update file data")
		render.Render(w, r, &errs.ErrorResponse{HttpStatus: http.StatusBadRequest, StatusText: "Cannot update file data"})
		return
	}

	reader, err := r.MultipartReader()
	if err != nil {
		log.WithError(err).Error("Could not parse multipart form")
		render.Render(w, r, errs.BadRequest(err))
		return
	}

	tempFile, err := ioutil.TempFile(s.Config.StorageDir, "temp_form_")
	if err != nil {
		log.WithError(err).Error("Could not create temp file")
		render.Render(w, r, errs.Internal(err))
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	buf := make([]byte, maxUploadBufferSize)
	hash := sha256.New()
	var fileSize int64 = 0

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.WithError(err).Error("Could not parse multipart form")
			render.Render(w, r, errs.BadRequest(err))
			return
		}

		if part.FormName() != "file" {
			log.WithField("formName", part.FormName()).Error("Unexpected form name")
			render.Render(w, r, errs.BadRequest(errors.New("Unexpected form name")))
			return
		}

		maxFileSize := s.Config.MaxFileSize
		fileWithHashReader := io.TeeReader(part, hash)
		reader := io.LimitReader(fileWithHashReader, maxFileSize+1-fileSize)

		written, err := io.CopyBuffer(tempFile, reader, buf)

		if err != nil {
			log.WithError(err).Error("Failed to write form data to temp file")
			render.Render(w, r, errs.Internal(err))
			return
		}

		fileSize += written

		if fileSize > maxFileSize {
			log.WithFields(log.Fields{
				"Total written": fileSize,
				"Max file size": maxFileSize,
			}).Error("Too large file size")
			render.Render(w, r, errs.EntityTooLarge)
			return
		}
	}

	if fileSize == 0 {
		log.WithError(err).Error("Could not find \"file\" form field")
		render.Render(w, r, errs.BadRequest(errors.New("No \"file\" form filed")))
		return
	}

	tempFile.Close()
	log.WithField("bytes", fileSize).Debug("Total written size")

	mime, err := mimetype.DetectFile(tempFile.Name())
	if err != nil {
		log.WithError(err).Error("Failed to detect mime type")
		file.ContentType = "application/octet-stream"
	} else {
		file.ContentType = mime.String()
	}

	fileHashValue := hex.EncodeToString(hash.Sum(nil))
	log.Debug("Got file with sha256 hash: %s", fileHashValue)

	filePath := filepath.Join(s.Config.StorageDir, fileHashValue)

	if fileNotExists(filePath) {
		err := os.Rename(tempFile.Name(), filePath)
		if err != nil {
			log.WithError(err).Error("Failed to move temp file to storage")
			render.Render(w, r, errs.Internal(err))
			return
		}
	}

	if _, err = s.Repo.IncFileHashRefCount(fileHashValue); err != nil {
		log.WithError(err).Error("Failed to ref count file")
		render.Render(w, r, errs.Internal(err))
		return
	}

	file.Hash = fileHashValue
	res, err := s.Repo.Update(*file)
	if err != nil {
		log.WithError(err).Error("Failed to update file")
		render.Render(w, r, errs.Internal(err))
		return
	}

	render.Respond(w, r, modelapi.FileResponse{File: &res})
}

// GetFile godoc
// @Summary Get existing file contents
// @Tags files
// @Produce json
// @Param id path string true "File id"
// @Success 200 {string} string
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /files/{id}/data [get]
func (s *FilesService) DownloadFileData(w http.ResponseWriter, r *http.Request) {
	file := s.getFile(r)
	if file == nil {
		log.Error("Failed to load file")
		render.Render(w, r, errs.NotFound)
		return
	}

	if file.Hash == "" {
		log.Error("Could not download file without blob")
		render.Render(w, r, errs.NotFound)
		return
	}

	filePath := filepath.Join(s.Config.StorageDir, file.Hash)
	stat, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		log.WithField("path", filePath).Error("Could not find requested file in the storage")
		render.Render(w, r, errs.Internal(errors.New("Could not find requested file")))
		return
	} else if err != nil {
		log.WithError(err).Error("os.Stat failed")
		render.Render(w, r, errs.Internal(err))
		return
	}

	reader, err := os.Open(filePath)
	if err != nil {
		log.WithError(err).Error("Could not open file")
		render.Render(w, r, errs.Internal(err))
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.Name))
	w.Header().Set("Content-Type", file.ContentType)
	w.Header().Set("Content-Length", strconv.FormatInt(stat.Size(), 10))

	bytesWritten, err := io.Copy(w, reader)

	if err != nil {
		log.WithError(err).Error("Failed to send file")
	} else if bytesWritten != stat.Size() {
		log.WithFields(log.Fields{
			"sentBytes": bytesWritten,
			"statBytes": stat.Size(),
		}).Error("Sent size does not match file size from stat")
	}
}
