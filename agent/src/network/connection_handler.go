package network

import (
	"context"
	"fmt"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
	"io"

	common "github.com/dc-lab/sky/agent/src/common"
	hardware "github.com/dc-lab/sky/agent/src/hardware"
	parser "github.com/dc-lab/sky/agent/src/parser"
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
			go SendHardwareData(client, hardware.GetTotalHardwareData(), hardware.GetFreeHardwareData())
		case *rm.TToAgentMessage_TaskRequest:
			fmt.Println("Task request")
			task := response.TaskRequest.GetTask()
			go StartTask(task)
		case *rm.TToAgentMessage_StageInRequest:
			fmt.Println("Stage in request")
			files := response.StageInRequest.GetFiles()
			taskId := response.StageInRequest.GetTaskId()
			go StageInFiles(client, taskId, files)
		case *rm.TToAgentMessage_StageOutRequest:
			fmt.Println("Stage out request")
			taskId := response.StageOutRequest.GetTaskId()
			go StageOutFiles(client, taskId)
		case *rm.TToAgentMessage_CancelTaskRequest:
			fmt.Println("Cancel task request")
			taskId := response.CancelTaskRequest.GetTaskId()
			go CancelTask(taskId)
		default:
			fmt.Println("Non type of response")
		}
	}
	// err = stream.CloseSend()
}

func RunClient() {
	stream, ctx := CreateConnection(parser.AgentConfig.ResourceManagerAddress)
	go SendRegistrationData(stream, parser.AgentConfig.Token)
	go ReceiveResourceManagerRequest(stream)
	go SendHealthChecks(stream)
	go UpdateTasksStatuses(stream)
	<-ctx.Done()
}
