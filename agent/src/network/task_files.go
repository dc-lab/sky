package network

import (
	utils "github.com/dc-lab/sky/agent/src/common"
	"github.com/dc-lab/sky/agent/src/data_manager"
	"github.com/dc-lab/sky/api/proto/common"
	"github.com/dc-lab/sky/api/proto/resource_manager"
	"io"
	"path"
)

func DownloadFiles(taskId string, files []*resource_manager.TFile) resource_manager.TStageInResponse {
	task, flag := GlobalTasksStatuses.Load(taskId)
	result := common.TResult{ResultCode: common.TResult_FAILED}
	if flag {
		for _, file := range files {
			err, body := data_manager_api.GetFileBody(file.GetId())
			if err != nil {
				result.ResultCode = common.TResult_FAILED
				err_str := err.Error()
				result.ErrorText = err_str
			}
			defer body.Close()
			out := utils.CreateFile(path.Join(task.ExecutionDir, file.GetAgentRelativeLocalPath()))
			defer out.Close()
			io.Copy(out, body)
		}
		if result.ResultCode != common.TResult_FAILED {
			result.ResultCode = common.TResult_SUCCESS
		}
	}
	return resource_manager.TStageInResponse{TaskId: taskId, Result: &result}
}

func StageInFiles(client resource_manager.ResourceManager_SendClient, taskId string, files []*resource_manager.TFile) {
	response := DownloadFiles(taskId, files)
	body := resource_manager.TFromAgentMessage_StageInResponse{StageInResponse: &response}
	err := client.Send(&resource_manager.TFromAgentMessage{Body: &body})
	utils.DealWithError(err)
}

func StageOutFiles(client resource_manager.ResourceManager_SendClient, taskId string) {
	task, flag := GlobalTasksStatuses.Load(taskId)
	if flag {
		filePaths := utils.GetChildrenFilePaths(task.ExecutionDir)
		var files []*resource_manager.TFile
		for _, filePath := range filePaths {
			file := data_manager_api.UploadFile(filePath, task.ExecutionDir)
			files = append(files, &file)
		}
		response := resource_manager.TStageOutResponse{TaskId: taskId, TaskFiles: files}
		body := resource_manager.TFromAgentMessage_StageOutResponse{StageOutResponse: &response}
		err := client.Send(&resource_manager.TFromAgentMessage{Body: &body})
		utils.DealWithError(err)
	}
}