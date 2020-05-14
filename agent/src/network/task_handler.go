package network

import (
	"fmt"
	"github.com/dc-lab/sky/agent/src/common"
	pb "github.com/dc-lab/sky/api/proto/common"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
	"os/exec"
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

func RunShellCommand(
	command string,
	directory string,
	stdOutFilePath string,
	stdErrFilePath string,
	beforeExecution func(pid int64, result *pb.TResult),
	afterExecution func(err error),
	quit <-chan struct{},
) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = directory
	stdoutFile := common.CreateFile(stdOutFilePath)
	defer stdoutFile.Close()
	cmd.Stdout = stdoutFile
	stderrFile := common.CreateFile(stdErrFilePath)
	defer stderrFile.Close()
	cmd.Stderr = stderrFile
	err := cmd.Start()
	common.DealWithError(err)
	pid := int64(cmd.Process.Pid)
	result := pb.TResult{ResultCode: pb.TResult_RUN}
	if beforeExecution != nil {
		beforeExecution(pid, &result)
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-quit:
		err := cmd.Process.Kill()
		common.DealWithError(err)
	case err := <-done:
		if afterExecution != nil {
			afterExecution(err)
		}
	}
}
