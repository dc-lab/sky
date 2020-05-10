package network

import (
	"context"
	"fmt"
	"io"

	common "github.com/dc-lab/sky/agent/src/common"
	hardware "github.com/dc-lab/sky/agent/src/hardware"
	parser "github.com/dc-lab/sky/agent/src/parser"
	pb "github.com/dc-lab/sky/agent/src/protos"
	"google.golang.org/grpc"
)

func CreateConnection(address string) (pb.ResourceManager_SendClient, context.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	common.DealWithError(err)

	client := pb.NewResourceManagerClient(conn)
	stream, err := client.Send(context.Background())
	common.DealWithError(err)
	ctx := stream.Context()
	common.DealWithError(err)
	return stream, ctx
}

func StageInFiles(client pb.ResourceManager_SendClient, task_id *string, files []*pb.TFile) {
	response := DownloadFiles(task_id, files)
	body := pb.TFromAgentMessage_StageInResponse{StageInResponse: &response}
	err := client.Send(&pb.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
}

func ReceiveResourceManagerRequest(client pb.ResourceManager_SendClient) {
	for {
		generalResponse, err := client.Recv()
		if err == io.EOF {
			return
		}
		common.DealWithError(err)
		switch response := generalResponse.Body.(type) {
		case *pb.TToAgentMessage_HardwareRequest:
			fmt.Println("Hardware data request")
			go SendHardwareData(client, hardware.GetHardwareData())
		case *pb.TToAgentMessage_TaskRequest:
			fmt.Println("Task request")
			task := response.TaskRequest.GetTask()
			go HandleTask(task)
		case *pb.TToAgentMessage_StageInRequest:
			fmt.Println("Stage in request")
			files := response.StageInRequest.Files
			task_id := response.StageInRequest.TaskId
			go StageInFiles(client, task_id, files)
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
