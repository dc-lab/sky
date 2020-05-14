package network

import (
	pb "github.com/dc-lab/sky/api/proto/common"
	"sync"
)

type TasksInfo struct {
	Mutex sync.RWMutex
	Data  map[string]*Task
}

func (info *TasksInfo) Load(key string) (*Task, bool) {
	info.Mutex.RLock()
	defer info.Mutex.RUnlock()
	val, ok := info.Data[key]
	return val, ok
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

func (info *TasksInfo) SetTaskResult(key string, result *pb.TResult) {
	info.Mutex.Lock()
	defer info.Mutex.Unlock()
	info.Data[key].Result = result
}

func (info *TasksInfo) GetTaskResult(key string) *pb.TResult {
	info.Mutex.RLock()
	defer info.Mutex.RUnlock()
	return info.Data[key].Result
}

var GlobalTasksStatuses = TasksInfo{
	Data: make(map[string]*Task),
}
