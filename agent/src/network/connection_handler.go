package network

import (
	"context"
	"fmt"
	"io"

	common "github.com/dc-lab/sky/common"
	hardware "github.com/dc-lab/sky/hardware"
	"github.com/dc-lab/sky/parser"
	pb "github.com/dc-lab/sky/protos"
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
			HandleTask(task)
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
