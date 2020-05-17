package network

import (
	"fmt"
	pb "github.com/dc-lab/sky/api/proto/common"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
)

func ConsumeTasksStatus(client rm.ResourceManager_SendClient, consumer func(rm.ResourceManager_SendClient, string, *pb.TResult)) {
	GlobalTasksStatuses.Mutex.RLock()
	defer GlobalTasksStatuses.Mutex.RUnlock()
	for taskID, processInfo := range GlobalTasksStatuses.Data {
		fmt.Println("key:", taskID, ", val:", processInfo)
		consumer(client, taskID, GlobalTasksStatuses.GetTaskResult(taskID)) // don't change last argument
	}
}

func StartTask(taskProto *rm.TTask) {
	result := pb.TResult{ResultCode: pb.TResult_WAIT}
	task := Task{Result: &result}
	task.Init(taskProto)
	GlobalTasksStatuses.Store(task.TaskId, &task)
	task.InstallRequirements()
	task.Run()
}

func CancelTask(taskId string) {
	task, flag := GlobalTasksStatuses.Load(taskId)
	if flag {
		task.Cancel()
	}
}

func DeleteTask(taskId string) {
	task, flag := GlobalTasksStatuses.Load(taskId)
	if flag {
		task.Delete()
	}
}
