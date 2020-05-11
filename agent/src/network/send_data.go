package network

import (
	"fmt"
	"time"

	common "github.com/dc-lab/sky/agent/src/common"
	hardware "github.com/dc-lab/sky/agent/src/hardware"
	pb "github.com/dc-lab/sky/api/proto/common"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
)

func printHardwareData(hwType string, hardwareData hardware.HardwareData) {
	fmt.Printf("%s %.2fc %db %db\n", hwType, hardwareData.CpuCount, hardwareData.MemoryBytes, hardwareData.DiskBytes)
}

func SendRegistrationData(client rm.ResourceManager_SendClient, token string) {
	request := rm.TGreetings{Token: token}
	body := rm.TFromAgentMessage_Greetings{Greetings: &request}
	err := client.Send(&rm.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
}

func SendHardwareData(client rm.ResourceManager_SendClient, totalHardwareData hardware.HardwareData, freeHardwareData hardware.HardwareData) {
	fmt.Println("Send hardware data")
	printHardwareData("total: ", totalHardwareData)
	printHardwareData("free: ", freeHardwareData)
	totalHardware := pb.THardwareData{CoresCount: totalHardwareData.CpuCount, MemoryBytes: totalHardwareData.MemoryBytes, DiskBytes: totalHardwareData.DiskBytes}
	freeHardware := pb.THardwareData{CoresCount: freeHardwareData.CpuCount, MemoryBytes: freeHardwareData.MemoryBytes, DiskBytes: freeHardwareData.DiskBytes}
	response := rm.THardwareResponse{TotalHardwareData: &totalHardware, FreeHardwareData: &freeHardware}
	body := rm.TFromAgentMessage_HardwareResponse{HardwareResponse: &response}
	err := client.Send(&rm.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
}

func SendTaskData(client rm.ResourceManager_SendClient, taskId string, resultPtr *pb.TResult) {
	fmt.Println("Send task data")
	request := rm.TTaskResponse{TaskId: taskId, Result: resultPtr}
	body := rm.TFromAgentMessage_TaskResponse{TaskResponse: &request}
	err := client.Send(&rm.TFromAgentMessage{Body: &body})
	common.DealWithError(err)
	if resultPtr.ResultCode == pb.TResult_FAILED || resultPtr.ResultCode == pb.TResult_SUCCESS {
		GlobalTasksStatuses.Delete(taskId)
	}
}

func SendHealthChecks(client rm.ResourceManager_SendClient) {
	for ; ; time.Sleep(time.Second * 10) {
		SendHardwareData(client, hardware.GetTotalHardwareData(), hardware.GetFreeHardwareData())
	}
}

func UpdateTasksStatuses(client rm.ResourceManager_SendClient) {
	for ; ; time.Sleep(time.Millisecond * 100) {
		ConsumeTasksStatus(client, SendTaskData)
	}
}
