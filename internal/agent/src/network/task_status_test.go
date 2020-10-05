package network

import (
	pb "github.com/dc-lab/sky/api/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ClearTasksStatuses() {
	GlobalTasksStatuses = TasksInfo{
		Data: make(map[string]*Task),
	}
}

func TestTaskInfoStoreAndLoad(t *testing.T) {
	ClearTasksStatuses()
	key := "key"
	_, flag := GlobalTasksStatuses.Load(key)
	assert.Equal(t, len(GlobalTasksStatuses.Data), 0)
	assert.Equal(t, flag, false)
	task := Task{TaskId: "custom_task_id"}
	GlobalTasksStatuses.Store(key, &task)
	assert.Equal(t, len(GlobalTasksStatuses.Data), 1)
	taskLoaded, flag := GlobalTasksStatuses.Load(key)
	assert.Equal(t, flag, true)
	assert.Equal(t, taskLoaded.TaskId, task.TaskId)
}

func TestTaskInfoDelete(t *testing.T) {
	ClearTasksStatuses()
	assert.Equal(t, len(GlobalTasksStatuses.Data), 0)
	key := "key"
	GlobalTasksStatuses.Store(key, nil)
	assert.Equal(t, len(GlobalTasksStatuses.Data), 1)
	GlobalTasksStatuses.Delete(key)
	assert.Equal(t, len(GlobalTasksStatuses.Data), 0)
}

func TestTaskInfoGetTaskResult(t *testing.T) {
	ClearTasksStatuses()
	key := "key"
	result := pb.TResult{ResultCode: pb.TResult_WAIT}
	task := Task{TaskId: "custom_task_id", Result: &result}
	GlobalTasksStatuses.Store(key, &task)
	resultLoaded := GlobalTasksStatuses.GetTaskResult(key)
	assert.Equal(t, resultLoaded.ResultCode, result.ResultCode)
}

func TestTaskInfoSetTaskResult(t *testing.T) {
	ClearTasksStatuses()
	key := "key"
	result := pb.TResult{ResultCode: pb.TResult_WAIT}
	task := Task{TaskId: "custom_task_id", Result: &result}
	GlobalTasksStatuses.Store(key, &task)
	result = pb.TResult{ResultCode: pb.TResult_SUCCESS}
	GlobalTasksStatuses.SetTaskResult(key, &result)
	resultLoaded := GlobalTasksStatuses.GetTaskResult(key)
	assert.Equal(t, resultLoaded.ResultCode, result.ResultCode)
}
