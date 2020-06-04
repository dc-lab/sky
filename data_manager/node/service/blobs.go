package service

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	pb "github.com/dc-lab/sky/api/proto/data_manager"
	"github.com/dc-lab/sky/data_manager/node/client"
	"github.com/dc-lab/sky/data_manager/node/config"
	errs "github.com/dc-lab/sky/data_manager/node/errors"
	"github.com/dc-lab/sky/data_manager/node/storage"
)

const maxUploadBufferSize = 8 * 1024 * 1024 // 8 MiB

type BlobsService struct {
	config *config.Config
	client *client.Client

	storage storage.Storage
}

func MakeBlobsService(config *config.Config, storage storage.Storage) (*BlobsService, error) {
	client, err := client.MakeClient(config)
	if err != nil {
		return nil, err
	}

	return &BlobsService{
		config,
		client,
		storage,
	}, nil
}

func (s *BlobsService) RunLoop() {
	ticker := time.NewTicker(s.config.PushInterval)
	for range ticker.C {
		log.Info("Start loop")
		status := &pb.NodeStatus{}
		s.storage.ListBlobs(func(hash string) error {
			status.BlobHashes = append(status.BlobHashes, hash)
			return nil
		})

		target, err := s.client.Loop(status)
		if err != nil {
			log.WithError(err).Errorln("Node loop failed")
			continue
		}

		log.Info("Got new target")
		for _, hash := range target.BlobHashes {
			log.Info("Hash: ", hash)
		}
		log.Info("Done")
	}
}

type OkResponse struct {
	Status string
}

func (s *BlobsService) DownloadBlob(w http.ResponseWriter, r *http.Request) {
	user_id := r.Header.Get("User-Id")
	if user_id == "" {
		log.Error("Unauthorized access")
		render.Render(w, r, errs.Unauthorized)
		return
	}

	file_id := chi.URLParam(r, "file_id")

	allow, hash, err := s.client.GetFileHash(file_id, user_id)
	if err != nil {
		log.WithFields(log.Fields{
			"user_id": user_id,
			"file_id": file_id,
		}).Error("Failed to get file hash")
		render.Render(w, r, errs.Internal(err))
		return
	}
	if !allow {
		log.WithFields(log.Fields{
			"user_id": user_id,
			"file_id": file_id,
			"error":   err,
		}).Error("File download forbidden")
		render.Render(w, r, errs.Forbidden)
		return
	}

	size, reader, err := s.storage.GetBlob(hash)
	if err != nil || reader == nil {
		log.WithFields(log.Fields{
			"user_id": user_id,
			"file_id": file_id,
			"hash":    hash,
			"error":   err,
		}).Error("Failed to get file from storage")
		render.Render(w, r, errs.Internal(errors.New("Failed to load file")))
		return
	}

	// FIXME(BigRedEye)
	// w.Header().Set("Content-Type", file.ContentType)
	// w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.Name))
	w.Header().Set("Content-Disposition", "attachment")
	w.Header().Set("Content-Length", strconv.FormatInt(size, 10))

	bytesWritten, err := io.Copy(w, reader)

	if err != nil {
		log.WithError(err).Error("Failed to send file")
	} else if bytesWritten != size {
		log.WithFields(log.Fields{
			"sentBytes": bytesWritten,
			"File size": size,
		}).Error("Sent size does not match file size")
	}
}

func (s *BlobsService) UploadBlob(w http.ResponseWriter, r *http.Request) {
	user_id := r.Header.Get("User-Id")
	if user_id == "" {
		log.Error("Unauthorized access")
		render.Render(w, r, errs.Unauthorized)
		return
	}

	file_id := chi.URLParam(r, "file_id")
	if file_id == "" {
		log.WithFields(log.Fields{
			"user_id": user_id,
			"file_id": file_id,
		}).Error("Failed to find file")
		render.Render(w, r, errs.NotFound)
		return
	}
	upload_token := r.URL.Query().Get("token")

	allow, err := s.client.ValidateUpload(user_id, file_id, upload_token)
	if err != nil {
		log.WithError(err).Error("Failed to validate upload")
		render.Render(w, r, errs.Internal(err))
		return
	}
	if !allow {
		log.WithError(err).Warning("File upload forbidden")
		render.Render(w, r, errs.Forbidden)
		return
	}

	reader, err := r.MultipartReader()
	if err != nil {
		log.WithError(err).Error("Failed to parse multipart form")
		render.Render(w, r, errs.BadRequest(err))
		return
	}

	writer, err := s.storage.PutBlob()
	if err != nil {
		log.WithError(err).Error("Failed to create temp file")
		render.Render(w, r, errs.Internal(err))
		return
	}

	defer writer.Discard()

	buf := make([]byte, maxUploadBufferSize)
	var bytesWritten int64 = 0

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.WithError(err).Error("Failed to parse multipart form")
			render.Render(w, r, errs.BadRequest(err))
			return
		}

		if part.FormName() != "file" {
			log.WithField("formName", part.FormName()).Error("Unexpected form name")
			render.Render(w, r, errs.BadRequest(errors.New("Unexpected form name")))
			return
		}

		maxFileSize := s.config.MaxFileSize
		reader := io.LimitReader(part, maxFileSize+1-bytesWritten)
		written, err := io.CopyBuffer(writer, reader, buf)

		if err != nil {
			log.WithError(err).Error("Failed to write form data to temp file")
			render.Render(w, r, errs.Internal(err))
			return
		}

		bytesWritten += written

		if bytesWritten > maxFileSize {
			log.WithFields(log.Fields{
				"Total written": bytesWritten,
				"Max file size": maxFileSize,
			}).Error("Too large file size")
			render.Render(w, r, errs.EntityTooLarge)
			return
		}
	}

	if bytesWritten == 0 {
		log.WithError(err).Error("Failed to find \"file\" form field")
		render.Render(w, r, errs.BadRequest(errors.New("No \"file\" form filed")))
		return
	}
	log.WithField("bytes", bytesWritten).Debug("Total file size")

	writer.Close()

	hash := writer.GetHash()
	log.Debug("Got file with sha256 hash: %s", hash)

	allow, err = s.client.SubmitFileHash(user_id, file_id, hash)
	if err != nil {
		log.WithError(err).Error("Failed to submit file hash")
		render.Render(w, r, errs.Internal(err))
		return
	}
	if !allow {
		log.WithError(err).Warning("File upload forbidden")
		render.Render(w, r, errs.Forbidden)
		return
	}

	err = writer.Finish()
	if err != nil {
		log.WithError(err).Error("Failed to finish file")
		render.Render(w, r, errs.Internal(err))
		return
	}

	render.Respond(w, r, OkResponse{Status: "File has been successfully uploaded"})
}
