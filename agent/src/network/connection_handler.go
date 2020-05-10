package network

import (
	"context"
	"fmt"
	data_manager_api "github.com/dc-lab/sky/agent/src/data_manager"
	"io"

	common "github.com/dc-lab/sky/agent/src/common"
	hardware "github.com/dc-lab/sky/agent/src/hardware"
	parser "github.com/dc-lab/sky/agent/src/parser"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
	"google.golang.org/grpc"
)

func CreateConnection(address string) (rm.ResourceManager_SendClient, context.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	common.DealWithError(err)

	client := rm.NewResourceManagerClient(conn)
	stream, err := client.Send(context.Background())
	common.DealWithError(err)
	ctx := stream.Context()
	common.DealWithError(err)
	return stream, ctx
}

func StageInFiles(client rm.ResourceManager_SendClient, task_id string, files []*rm.TFile) {
	response := DownloadFiles(task_id, files)
	body := rm.TFromAgentMessage_StageInResponse{StageInResponse: &response}
	err := client.Send(&rm.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
}

func StageOutFiles(client pb.ResourceManager_SendClient, task_id string) {
	taskDir := GetExecutionDirForTaskId(task_id)
	filePaths := common.GetChildrenFilePaths(taskDir)
	var files []*pb.TFile
	for _, filePath := range filePaths {
		file := data_manager_api.UploadFile(filePath, taskDir)
		files = append(files, &file)
	}
	response := pb.TStageOutResponse{TaskId: &task_id, TaskFiles: files}
	body := pb.TFromAgentMessage_StageOutResponse{StageOutResponse: &response}
	err := client.Send(&pb.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
}

func ReceiveResourceManagerRequest(client rm.ResourceManager_SendClient) {
	for {
		generalResponse, err := client.Recv()
		if err == io.EOF {
			return
		}
		common.DealWithError(err)
		switch response := generalResponse.Body.(type) {
		case *rm.TToAgentMessage_HardwareRequest:
			fmt.Println("Hardware data request")
			go SendHardwareData(client, hardware.GetHardwareData())
		case *rm.TToAgentMessage_TaskRequest:
			fmt.Println("Task request")
			task := response.TaskRequest.GetTask()
			go HandleTask(task)
		case *rm.TToAgentMessage_StageInRequest:
			fmt.Println("Stage in request")
			files := response.StageInRequest.GetFiles()
			taskId := response.StageInRequest.GetTaskId()
			go StageInFiles(client, taskId, files)
		case *rm.TToAgentMessage_SageOuRequest:
			fmt.Println("Stage out request")
			taskId := response.SageOuRequest.GetTaskId()
			go StageOutFiles(client, taskId)
		default:
			fmt.Println("Non type of response")
		}
	}
	// err = stream.CloseSend()
}

func RunClient() {
	stream, ctx := CreateConnection(parser.AgentConfig.ResourceManagerAddress)
	go SendRegistrationData(stream, &parser.AgentConfig.Token)
	go ReceiveResourceManagerRequest(stream)
	go SendHealthChecks(stream)
	go UpdateTasksStatuses(stream)
	<-ctx.Done()
}
