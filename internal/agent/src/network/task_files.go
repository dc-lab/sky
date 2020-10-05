package network

import (
	"io"
	"path"

	pb "github.com/dc-lab/sky/api/proto"
	utils "github.com/dc-lab/sky/internal/agent/src/common"
	"github.com/dc-lab/sky/internal/agent/src/data_manager"
	"github.com/dc-lab/sky/internal/agent/src/local_cache"
	"github.com/dc-lab/sky/internal/agent/src/parser"
)

type FileStorage struct {
	config *parser.Config
	cache  *local_cache.Cache
}

func NewFileStorage(config *parser.Config) (*FileStorage, error) {
	cache, err := local_cache.NewCache(config)
	if err != nil {
		return nil, err
	}

	return &FileStorage{
		config: config,
		cache:  cache,
	}, nil
}

func (s *FileStorage) DownloadFile(localPath string, file *pb.TFile) error {
	reader, err := s.cache.GetCacheFileReader(file.GetHash())
	if err != nil {
		err, reader = data_manager_api.GetFileBody(file.GetId())
		if err == nil {
			out, err := s.cache.NewCacheFileWriter(file.GetHash())
			if err == nil {
				_, err = io.Copy(out, reader)
			}
		}
		reader, err = s.cache.GetCacheFileReader(file.GetHash())
	}
	if err == nil {
		out := utils.CreateFile(localPath)
		_, err = io.Copy(out, reader)
	}
	return err
}

func (s *FileStorage) DownloadFiles(taskId string, files []*pb.TFile) pb.TStageInResponse {
	taskExecutionDir, err := GetTaskExecutionDir(taskId, s.config.AgentDirectory)
	utils.DealWithError(err)
	result := pb.TResult{ResultCode: pb.TResult_FAILED}
	for _, file := range files {
		localPath := path.Join(taskExecutionDir, file.GetAgentRelativeLocalPath())
		err := s.DownloadFile(localPath, file)
		if err != nil {
			result.ResultCode = pb.TResult_FAILED
			err_str := err.Error()
			result.ErrorText = err_str
		}
	}
	if result.ResultCode != pb.TResult_FAILED {
		result.ResultCode = pb.TResult_SUCCESS
	}
	return pb.TStageInResponse{TaskId: taskId, Result: &result}
}

func (s *FileStorage) StageInFiles(client pb.ResourceManager_SendClient, taskId string, files []*pb.TFile) {
	response := s.DownloadFiles(taskId, files)
	body := pb.TFromAgentMessage_StageInResponse{StageInResponse: &response}
	err := client.Send(&pb.TFromAgentMessage{Body: &body})
	utils.DealWithError(err)
}

func (s *FileStorage) UploadTaskFile(taskId string, filePath string) *pb.TFile {
	task, flag := GlobalTasksStatuses.Load(taskId)
	var file pb.TFile
	if flag {
		file = data_manager_api.UploadFile(filePath, task.ExecutionDir)
	}
	return &file
}

func (s *FileStorage) StageOutFiles(client pb.ResourceManager_SendClient, taskId string, localPath string) {
	file := s.UploadTaskFile(taskId, localPath)
	response := pb.TStageOutResponse{TaskId: taskId, TaskFile: file}
	body := pb.TFromAgentMessage_StageOutResponse{StageOutResponse: &response}
	err := client.Send(&pb.TFromAgentMessage{Body: &body})
	utils.DealWithError(err)
}
