package network

import (
	"fmt"

	pb "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/agent/src/parser"
)

func ConsumeTasksData(client pb.ResourceManager_SendClient, consumer func(pb.ResourceManager_SendClient, string, *pb.Result)) {
	GlobalTasksStatuses.Mutex.RLock()
	defer GlobalTasksStatuses.Mutex.RUnlock()
	for taskID, processInfo := range GlobalTasksStatuses.Data {
		fmt.Println("key:", taskID, ", val:", processInfo)
		consumer(client, taskID, GlobalTasksStatuses.GetTaskResult(taskID)) // don't change last argument
	}
}

func StartTask(taskProto *pb.Task, config *parser.Config) {
	result := pb.Result{ResultCode: pb.Result_WAIT}
	task := Task{Result: &result}
	task.Init(taskProto, config)
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
