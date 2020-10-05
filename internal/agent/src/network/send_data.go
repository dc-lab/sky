package network

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"

	pb "github.com/dc-lab/sky/api/proto"
	common "github.com/dc-lab/sky/internal/agent/src/common"
	hardware "github.com/dc-lab/sky/internal/agent/src/hardware"
)

func printHardwareData(hwType string, hardwareData hardware.HardwareData) {
	log.Printf("%s %.2fc %db %db\n", hwType, hardwareData.CpuCount, hardwareData.MemoryBytes, hardwareData.DiskBytes)
}

func ResourceRegistration(client pb.ResourceManager_SendClient, token string) bool {
	request := pb.Greetings{Token: token}
	body := pb.FromAgentMessage_Greetings{Greetings: &request}
	err := client.Send(&pb.FromAgentMessage{Body: &body})
	common.DealWithError(err)
	generalResponse, err := client.Recv()
	if err == io.EOF {
		return false
	}
	common.DealWithError(err)
	switch response := generalResponse.Body.(type) {
	case *pb.ToAgentMessage_GreetingsValidation:
		return response.GreetingsValidation.GetResult().GetResultCode() == pb.Result_SUCCESS
	default:
		return false
	}
}

func SendHardwareData(client pb.ResourceManager_SendClient, totalHardwareData hardware.HardwareData, freeHardwareData hardware.HardwareData) {
	log.Println("Send hardware data")
	printHardwareData("total: ", totalHardwareData)
	printHardwareData("free: ", freeHardwareData)
	totalHardware := pb.HardwareData{CoresCount: totalHardwareData.CpuCount, MemoryBytes: totalHardwareData.MemoryBytes, DiskBytes: totalHardwareData.DiskBytes}
	freeHardware := pb.HardwareData{CoresCount: freeHardwareData.CpuCount, MemoryBytes: freeHardwareData.MemoryBytes, DiskBytes: freeHardwareData.DiskBytes}
	response := pb.HardwareResponse{TotalHardwareData: &totalHardware, FreeHardwareData: &freeHardware}
	body := pb.FromAgentMessage_HardwareResponse{HardwareResponse: &response}
	err := client.Send(&pb.FromAgentMessage{Body: &body})
	common.DealWithError(err)
}

func gatherTaskResults(root string) ([]*pb.TaskFile, error) {
	var files []*pb.TaskFile
	err := common.ListChildrenFiles(root, func(path string, info os.FileInfo) error {
		mtime, err := ptypes.TimestampProto(info.ModTime())
		if err != nil {
			return err
		}
		// TODO Calculate hash here?
		files = append(files, &pb.TaskFile{
			Path:             path,
			Size:             info.Size(),
			ModificationTime: mtime,
		})
		return nil
	})

	return files, err
}

func SendTaskData(client pb.ResourceManager_SendClient, taskId string, resultPtr *pb.Result) {
	task, _ := GlobalTasksStatuses.Load(taskId)
	files, err := gatherTaskResults(task.ExecutionDir)
	common.DealWithError(err)
	fmt.Println("Send task data")
	request := pb.TaskResponse{TaskId: taskId, Result: resultPtr, TaskFiles: files}
	body := pb.FromAgentMessage_TaskResponse{TaskResponse: &request}
	err = client.Send(&pb.FromAgentMessage{Body: &body})
	common.DealWithError(err)
	//if resultPtr.ResultCode == pb.Result_FAILED || resultPtr.ResultCode == pb.Result_SUCCESS {
	//	GlobalTasksStatuses.Delete(taskId)
	//}
}

func UpdateTasksInfo(client pb.ResourceManager_SendClient) {
	for ; ; time.Sleep(time.Millisecond * 100) {
		ConsumeTasksData(client, SendTaskData)
	}
}
