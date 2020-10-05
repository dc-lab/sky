package service

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/dc-lab/sky/internal/data_manager/master/config"
	"github.com/dc-lab/sky/internal/data_manager/master/modeldb"
	"github.com/dc-lab/sky/internal/data_manager/master/repo"
)

type FilesService struct {
	Repo   *repo.FilesRepo
	Config *config.Config
}

func (s *FilesService) CreateFile(req *modeldb.File) (*modeldb.File, error) {
	file, err := s.Repo.Create(req)
	if err != nil {
		log.WithError(err).Error("Failed to insert new file into db")
		return nil, err
	}

	uploadLocations, err := s.Repo.SelectBestUploadStorages()
	if err != nil {
		log.WithError(err).Error("Failed to select best upload storages")
		return nil, err
	}

	file.UploadUrls = make([]string, 0)
	for _, location := range uploadLocations {
		file.UploadUrls = append(file.UploadUrls, makeFileUploadUrl(location, file.Id, file.UploadToken))
	}
	return file, err
}

func makeFileUploadUrl(location string, id string, uploadToken string) string {
	return makeFileDownloadUrl(location, id) + fmt.Sprintf("?token=%s", uploadToken)
}

func makeFileDownloadUrl(location string, id string) string {
	return fmt.Sprintf("http://%s/v1/files/%s/data", location, id)
}

func (s *FilesService) GetFile(id string) (*modeldb.File, error) {
	file, err := s.Repo.Get(id)
	if err != nil {
		log.WithError(err).Error("Failed to get file from db")
		return nil, err
	}
	return file, err
}

func (s *FilesService) GetFileLocation(id string) ([]string, error) {
	locations, err := s.Repo.GetFileLocations(id)
	if err != nil {
		log.WithError(err).Error("Failed to get file from db")
		return nil, err
	}

	urls := make([]string, 0)
	for _, loc := range locations {
		urls = append(urls, makeFileDownloadUrl(loc, id))
	}
	return urls, err
}

func (s *FilesService) ListTaskResults(task_id string, path_prefix string) ([]modeldb.File, error) {
	files, err := s.Repo.GetTaskResults(task_id, path_prefix)

	if err != nil {
		log.WithFields(log.Fields{
			"task_id":     task_id,
			"path_prefix": path_prefix,
		}).WithError(err).Error("Failed to get task results from db")
	}
	return files, err
}

func (s *FilesService) TrySetFileHash(userId string, nodeId string, fileId string, hash string) (bool, error) {
	file, _ := s.Repo.Get(fileId)
	log.WithFields(log.Fields{
		"fileId": fileId,
		"file":   file.Hash,
	}).Warn("Start set hash")
	ok, err := s.Repo.SetFileHash(fileId, hash)
	if err != nil {
		log.WithFields(log.Fields{
			"fileId": fileId,
		}).WithError(err).Error("Failed to set file hash")
		return false, err
	}
	log.WithFields(log.Fields{
		"fileId": fileId,
	}).Warn("Finish set hash")

	if !ok {
		return ok, nil
	}

	_, err = s.Repo.IncFileHashRefCount(hash)
	if err != nil {
		log.WithFields(log.Fields{
			"fileId": fileId,
			"hash":   hash,
		}).WithError(err).Error("Failed to update hash count")
		return false, err
	}

	err = s.Repo.AddHashTarget(hash, nodeId)
	if err != nil {
		log.WithFields(log.Fields{
			"fileId": fileId,
			"hash":   hash,
		}).WithError(err).Error("Failed to set hash target")
		return false, err
	}

	return ok, err
}

func (s *FilesService) ProcessNodeReport(nodeId string, freeSpace int64, blobs []string) ([]string, error) {
	log.Infof("Report from node: %s", nodeId)

	err := s.Repo.UpdateNodeReportTimestamp(nodeId, freeSpace)
	if err != nil {
		log.WithFields(log.Fields{
			"nodeI": nodeId,
		}).WithError(err).Error("Failed to update node report timeout")
		return nil, err
	}

	err = s.Repo.UpdateLocations(nodeId, blobs)
	if err != nil {
		log.WithFields(log.Fields{
			"nodeId": nodeId,
		}).WithError(err).Error("Failed to update blobs locations from node report")
		return nil, err
	}

	return s.Repo.GetBlobsForLocation(nodeId)
}
