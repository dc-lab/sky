package network

import (
	"sync"

	pb "github.com/dc-lab/sky/protos"
)

type ProcessInfo struct {
	Result    pb.TResult
	ProcessID int
}

type TasksInfo struct {
	Mutex sync.RWMutex
	Data  map[string]ProcessInfo
}

func (info *TasksInfo) Load(key string) (ProcessInfo, bool) {
	info.Mutex.RLock()
	defer info.Mutex.RUnlock()
	val, ok := info.Data[key]
	return val, ok
}

func (info *TasksInfo) Store(key string, value ProcessInfo) {
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

var GlobalTasksStatuses = TasksInfo{
	Data: make(map[string]ProcessInfo),
}
