package network

import (
	pb "github.com/dc-lab/sky/api/proto/common"
	"sync"
)

type TasksInfo struct {
	Mutex sync.RWMutex
	Data  map[string]*Task
}

func (info *TasksInfo) Load(key string) (Task, bool) {
	info.Mutex.RLock()
	defer info.Mutex.RUnlock()
	val, ok := info.Data[key]
	return *val, ok
}

func (info *TasksInfo) Store(key string, value *Task) {
	info.Mutex.Lock()
	defer info.Mutex.Unlock()
	info.Data[key] = value
}

func (info *TasksInfo) Delete(key string) {
	_, ok := info.Load(key)
	info.Mutex.Lock()
	defer info.Mutex.Unlock()
	if ok {
		delete(info.Data, key)
	}
}

func (info *TasksInfo) UpdateTaskResult(taskId string, result *pb.TResult) {
	task, _ := GlobalTasksStatuses.Load(taskId)
	task.Result = result
	GlobalTasksStatuses.Store(taskId, &task)
}

// TODO: delete
func (info *TasksInfo) UpdateTaskProcessID(taskId string, pid int64) {
	task, _ := GlobalTasksStatuses.Load(taskId)
	task.ProcessID = pid
	GlobalTasksStatuses.Store(taskId, &task)
}

var GlobalTasksStatuses = TasksInfo{
	Data: make(map[string]*Task),
}
