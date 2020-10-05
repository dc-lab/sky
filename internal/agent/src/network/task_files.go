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

func (s *FileStorage) DownloadFile(localPath string, file *pb.FileOnAgent) error {
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

func (s *FileStorage) DownloadFiles(taskId string, files []*pb.FileOnAgent) pb.StageInResponse {
	taskExecutionDir, err := GetTaskExecutionDir(taskId, s.config.AgentDirectory)
	utils.DealWithError(err)
	result := pb.Result{ResultCode: pb.Result_FAILED}
	for _, file := range files {
		localPath := path.Join(taskExecutionDir, file.GetAgentRelativeLocalPath())
		err := s.DownloadFile(localPath, file)
		if err != nil {
			result.ResultCode = pb.Result_FAILED
			err_str := err.Error()
			result.ErrorText = err_str
		}
	}
	if result.ResultCode != pb.Result_FAILED {
		result.ResultCode = pb.Result_SUCCESS
	}
	return pb.StageInResponse{TaskId: taskId, Result: &result}
}

func (s *FileStorage) StageInFiles(client pb.ResourceManager_SendClient, taskId string, files []*pb.FileOnAgent) {
	response := s.DownloadFiles(taskId, files)
	body := pb.FromAgentMessage_StageInResponse{StageInResponse: &response}
	err := client.Send(&pb.FromAgentMessage{Body: &body})
	utils.DealWithError(err)
}

func (s *FileStorage) UploadTaskFile(taskId string, filePath string) *pb.FileOnAgent {
	task, flag := GlobalTasksStatuses.Load(taskId)
	var file pb.FileOnAgent
	if flag {
		file = data_manager_api.UploadFile(filePath, task.ExecutionDir)
	}
	return &file
}

func (s *FileStorage) StageOutFiles(client pb.ResourceManager_SendClient, taskId string, localPath string) {
	file := s.UploadTaskFile(taskId, localPath)
	response := pb.StageOutResponse{TaskId: taskId, TaskFile: file}
	body := pb.FromAgentMessage_StageOutResponse{StageOutResponse: &response}
	err := client.Send(&pb.FromAgentMessage{Body: &body})
	utils.DealWithError(err)
}
