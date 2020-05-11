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
		consumer(client, taskID, processInfo.Result) // don't change last argument
	}
}

func StartTask(taskProto *rm.TTask) {
	result := pb.TResult{ResultCode: pb.TResult_WAIT}
	task := Task{Result: &result}
	task.Init(taskProto)
	GlobalTasksStatuses.Store(task.TaskId, &task)
	task.InstallRequirements()
	task.Run()
	//RunShellCommand2(
	//	task.RequirementsShellCommand,
	//	task.ExecutionDir,
	//	path.Join(task.ExecutionDir, "requirements_out"),
	//	path.Join(task.ExecutionDir, "requirements_err"),
	//	task.TaskId,
	//	false)
	//RunShellCommand2(
	//	task.ExecutionShellCommand,
	//	task.ExecutionDir,
	//	path.Join(task.ExecutionDir, "execution_out"),
	//	path.Join(task.ExecutionDir, "execution_err"),
	//	task.TaskId,
	//	true)
}

func CancelTask(taskId string) {
	task, flag := GlobalTasksStatuses.Load(taskId)
	if flag {
		task.Cancel()
	}
}

func RunShellCommand2(command string, directory string, stdOutFilePath string, stdErrFilePath string, taskId string, changeTaskStatus bool) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = directory
	stdoutFile := common.CreateFile(stdOutFilePath)
	defer stdoutFile.Close()
	cmd.Stdout = stdoutFile
	stderrFile := common.CreateFile(stdErrFilePath)
	defer stderrFile.Close()
	cmd.Stderr = stderrFile
	pid := int64(cmd.Process.Pid)
	result := pb.TResult{ResultCode: pb.TResult_RUN}
	if changeTaskStatus {
		GlobalTasksStatuses.UpdateTaskResult(taskId, &result)
		GlobalTasksStatuses.UpdateTaskProcessID(taskId, pid)
	}
	err := cmd.Run()
	if changeTaskStatus {
	} else {
		common.DealWithError(err)
	}
}

func RunShellCommand(
	command string,
	directory string,
	stdOutFilePath string,
	stdErrFilePath string,
	beforeExecution func(pid int64, result *pb.TResult),
	afterExecution func(err error),
) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = directory
	stdoutFile := common.CreateFile(stdOutFilePath)
	defer stdoutFile.Close()
	cmd.Stdout = stdoutFile
	stderrFile := common.CreateFile(stdErrFilePath)
	defer stderrFile.Close()
	cmd.Stderr = stderrFile
	cmd.Start()
	pid := int64(cmd.Process.Pid)
	result := pb.TResult{ResultCode: pb.TResult_RUN}
	if beforeExecution != nil {
		beforeExecution(pid, &result)
	}
	err := cmd.Wait()
	//err := cmd.Run()
	if afterExecution != nil {
		afterExecution(err)
	}
}
