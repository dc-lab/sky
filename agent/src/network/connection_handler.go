package network

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

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
			localPath := response.StageOutRequest.GetAgentRelativeLocalPath()
			go StageOutFiles(client, taskId, localPath)
		case *rm.TToAgentMessage_CancelTaskRequest:
			taskId := response.CancelTaskRequest.GetTaskId()
			go CancelTask(taskId)
		case *rm.TToAgentMessage_DeleteTaskRequest:
			taskId := response.DeleteTaskRequest.GetTaskId()
			go DeleteTask(taskId)
		default:
			fmt.Println("Non type of response")
		}
	}
	// err = stream.CloseSend()
}

func UpdateHealthFile(healthFilePath string) {
	for ; ; time.Sleep(time.Millisecond * 200) {
		tsString := common.CurrentTimestampMillisString()
		err := ioutil.WriteFile(healthFilePath, []byte(tsString), 0644)
		common.DieWithError(err)
	}
}

func RunClient() {
	stream, ctx := CreateConnection(parser.AgentConfig.ResourceManagerAddress)
	success := ResourceRegistration(stream, parser.AgentConfig.Token)
	if !success {
		log.Println("Failed resource registration. Invalid greetings")
		os.Exit(1)
	}
	go ReceiveResourceManagerRequest(stream)
	go SendHealthChecks(stream)
	go UpdateTasksInfo(stream)
	go UpdateHealthFile(parser.AgentConfig.HealthFile)
	<-ctx.Done()
}
