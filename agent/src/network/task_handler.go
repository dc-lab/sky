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
	pb "github.com/dc-lab/sky/agent/src/protos"
)

func ConsumeTasksStatus(client pb.ResourceManager_SendClient, consumer func(pb.ResourceManager_SendClient, string, *pb.TResult)) {
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

func HandleTask(task *pb.TTask) {
	result := pb.TResult{ResultCode: pb.TResult_WAIT.Enum()}
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

func StageInFiles(task_id *string, files []*pb.TFile) {
	executionDir := GetExecutionDirForTaskId(*task_id)
	err := os.Mkdir(executionDir, 0777)
	common.DealWithError(err)
	for _, file := range files {
		body := data_manager_api.GetFileBody(file.GetId())
		defer body.Close()
		out := common.CreateFile(path.Join(executionDir, file.GetAgentRelativeLocalPath()))
		defer out.Close()
		io.Copy(out, body)
	}
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
	result := pb.TResult{ResultCode: pb.TResult_RUN.Enum()}
	if changeTaskStatus {
		GlobalTasksStatuses.Store(taskId, ProcessInfo{Result: result})
	}
	err := cmd.Run()
	if changeTaskStatus {
		if err != nil {
			result = pb.TResult{ResultCode: pb.TResult_FAILED.Enum(), ErrorCode: pb.TResult_INTERNAL.Enum()}
			fmt.Println("Error ", err)
			common.DealWithError(err)
		} else {
			result = pb.TResult{ResultCode: pb.TResult_SUCCESS.Enum()}
		}
		GlobalTasksStatuses.Store(taskId, ProcessInfo{Result: result})
	} else {
		common.DealWithError(err)
	}
}
