package network

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/dc-lab/sky/agent/src/common"
	data_manager_api "github.com/dc-lab/sky/agent/src/data_manager"
	parser "github.com/dc-lab/sky/agent/src/parser"
	pb "github.com/dc-lab/sky/api/proto/common"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
)

func ConsumeTasksStatus(client rm.ResourceManager_SendClient, consumer func(rm.ResourceManager_SendClient, string, *pb.TResult)) {
	GlobalTasksStatuses.Mutex.RLock()
	defer GlobalTasksStatuses.Mutex.RUnlock()
	for taskID, processInfo := range GlobalTasksStatuses.Data {
		fmt.Println("key:", taskID, ", val:", processInfo)
		consumer(client, taskID, &processInfo.Result) // don't change last argument
	}
}

func GetExecutionDirForTaskId(task_id string) string {
	return path.Join(parser.AgentConfig.AgentDirectory, task_id)
}

func HandleTask(task *rm.TTask) {
	result := pb.TResult{ResultCode: pb.TResult_WAIT}
	GlobalTasksStatuses.Store(task.GetId(), ProcessInfo{Result: result})
	executionDir := GetExecutionDirForTaskId(task.GetId())
	err := os.Mkdir(executionDir, 0777)
	common.DealWithError(err)
	RunShellCommand(
		task.GetRequirementsShellCommand(),
		executionDir,
		path.Join(executionDir, "requirements_out"),
		path.Join(executionDir, "requirements_err"),
		task.GetId(),
		false)
	RunShellCommand(
		task.GetExecutionShellCommand(),
		executionDir,
		path.Join(executionDir, "execution_out"),
		path.Join(executionDir, "execution_err"),
		task.GetId(),
		true)
}

func DownloadFiles(task_id string, files []*rm.TFile) rm.TStageInResponse {
	executionDir := GetExecutionDirForTaskId(task_id)
	err := os.Mkdir(executionDir, 0777)
	common.DealWithError(err)
	result := pb.TResult{ResultCode: pb.TResult_RUN}
	for _, file := range files {
		err, body := data_manager_api.GetFileBody(file.GetId())
		if err != nil {
			result.ResultCode = pb.TResult_FAILED
			err_str := err.Error()
			result.ErrorText = err_str
		}
		defer body.Close()
		out := common.CreateFile(path.Join(executionDir, file.GetAgentRelativeLocalPath()))
		defer out.Close()
		io.Copy(out, body)
	}
	if result.ResultCode != pb.TResult_FAILED {
		result.ResultCode = pb.TResult_SUCCESS
	}
	return rm.TStageInResponse{TaskId: &task_id, Result: &result}
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
		GlobalTasksStatuses.Store(taskId, ProcessInfo{Result: result})
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
		GlobalTasksStatuses.Store(taskId, ProcessInfo{Result: result})
	} else {
		common.DealWithError(err)
	}
}
