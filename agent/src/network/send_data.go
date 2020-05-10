package network

import (
	"fmt"
	"time"

	common "github.com/dc-lab/sky/agent/src/common"
	hardware "github.com/dc-lab/sky/agent/src/hardware"
	pb "github.com/dc-lab/sky/api/proto/common"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
)

func SendRegistrationData(client rm.ResourceManager_SendClient, token *string) {
	request := rm.TGreetings{Token: token}
	body := rm.TFromAgentMessage_Greetings{Greetings: &request}
	err := client.Send(&rm.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
}

func SendHardwareData(client rm.ResourceManager_SendClient, hardwareData hardware.HardwareData) {
	fmt.Println("Send hardware data")
	request := pb.THardwareData{CoresCount: hardwareData.CpuCount, MemoryBytes: hardwareData.MemoryBytes, DiskBytes: hardwareData.DiskBytes}
	body := rm.TFromAgentMessage_HardwareData{HardwareData: &request}
	err := client.Send(&rm.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
}

func SendTaskData(client rm.ResourceManager_SendClient, taskId string, resultPtr *pb.TResult) {
	fmt.Println("Send task data")
	request := rm.TTaskResponse{TaskId: &taskId, Result: resultPtr}
	body := rm.TFromAgentMessage_TaskResponse{TaskResponse: &request}
	err := client.Send(&rm.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
	if resultPtr.ResultCode == pb.TResult_FAILED || resultPtr.ResultCode == pb.TResult_SUCCESS {
		GlobalTasksStatuses.Delete(taskId)
	}
}

func SendHealthChecks(client rm.ResourceManager_SendClient) {
	for ; ; time.Sleep(time.Second * 10) {
		SendHardwareData(client, hardware.GetHardwareData())
	}
}

func UpdateTasksStatuses(client rm.ResourceManager_SendClient) {
	for ; ; time.Sleep(time.Microsecond * 100) {
		ConsumeTasksStatus(client, SendTaskData)
	}
}
