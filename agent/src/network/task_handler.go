package network

import (
	"fmt"
	"os/exec"
	"path"

	"github.com/dc-lab/sky/agent/src/common"
	pb "github.com/dc-lab/sky/api/proto/common"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
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
	task := Task{TaskId: taskProto.GetId(), Result: &result}
	task.Init()
	GlobalTasksStatuses.Store(taskProto.GetId(), &task)
	RunShellCommand(
		taskProto.GetRequirementsShellCommand(),
		task.ExecutionDir,
		path.Join(task.ExecutionDir, "requirements_out"),
		path.Join(task.ExecutionDir, "requirements_err"),
		taskProto.GetId(),
		false)
	RunShellCommand(
		taskProto.GetExecutionShellCommand(),
		task.ExecutionDir,
		path.Join(task.ExecutionDir, "execution_out"),
		path.Join(task.ExecutionDir, "execution_err"),
		taskProto.GetId(),
		true)
}

func RunShellCommand(command string, directory string, stdOutFilePath string, stdErrFilePath string, taskId string, changeTaskStatus bool) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = directory
	fmt.Println("Directory ", directory)
	fmt.Println("Command ", command)
	stdoutFile := common.CreateFile(stdOutFilePath)
	defer stdoutFile.Close()
	cmd.Stdout = stdoutFile
	stderrFile := common.CreateFile(stdErrFilePath)
	defer stderrFile.Close()
	cmd.Stderr = stderrFile
	// pid := cmd.Process.Pid
	result := pb.TResult{ResultCode: pb.TResult_RUN}
	if changeTaskStatus {
		GlobalTasksStatuses.UpdateTaskResult(taskId, &result)
	}
	err := cmd.Run()
	if changeTaskStatus {
		if err != nil {
			result = pb.TResult{ResultCode: pb.TResult_FAILED, ErrorCode: pb.TResult_INTERNAL}
			fmt.Println("Error ", err)
			common.DealWithError(err)
		} else {
			result = pb.TResult{ResultCode: pb.TResult_SUCCESS}
		}
		GlobalTasksStatuses.UpdateTaskResult(taskId, &result)
	} else {
		common.DealWithError(err)
	}
}
