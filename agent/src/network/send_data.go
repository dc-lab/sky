package network

import (
	"fmt"
	"time"

	common "github.com/dc-lab/sky/agent/src/common"
	hardware "github.com/dc-lab/sky/agent/src/hardware"
	pb "github.com/dc-lab/sky/agent/src/protos"
)

func SendRegistrationData(client pb.ResourceManager_SendClient, token *string) {
	request := pb.TGreetings{Token: token}
	body := pb.TFromAgentMessage_Greetings{Greetings: &request}
	err := client.Send(&pb.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
}

func SendHardwareData(client pb.ResourceManager_SendClient, hardwareData hardware.HardwareData) {
	fmt.Println("Send hardware data")
	request := pb.THardwareData{CoresCount: &hardwareData.CpuCount, MemoryAmount: &hardwareData.MemoryAmount, DiskAmount: &hardwareData.DiskAmount}
	body := pb.TFromAgentMessage_HardwareData{HardwareData: &request}
	err := client.Send(&pb.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
}

func SendTaskData(client pb.ResourceManager_SendClient, taskId string, resultPtr *pb.TResult) {
	fmt.Println("Send task data")
	request := pb.TTaskResponse{TaskId: &taskId, Result: resultPtr}
	body := pb.TFromAgentMessage_TaskResponse{TaskResponse: &request}
	err := client.Send(&pb.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
	if *resultPtr.ResultCode == pb.TResult_FAILED || *resultPtr.ResultCode == pb.TResult_SUCCESS {
		GlobalTasksStatuses.Delete(taskId)
	}
}

func SendHealthChecks(client pb.ResourceManager_SendClient) {
	for ; ; time.Sleep(time.Second * 10) {
		SendHardwareData(client, hardware.GetHardwareData())
	}
}

func UpdateTasksStatuses(client pb.ResourceManager_SendClient) {
	for ; ; time.Sleep(time.Microsecond * 100) {
		ConsumeTasksStatus(client, SendTaskData)
	}
}
