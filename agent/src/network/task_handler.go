package network

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/dc-lab/sky/common"
	parser "github.com/dc-lab/sky/parser"
	pb "github.com/dc-lab/sky/protos"
)

func ConsumeTasksStatus(client pb.ResourceManager_SendClient, consumer func(pb.ResourceManager_SendClient, string, *pb.TResult)) {
	GlobalTasksStatuses.Mutex.RLock()
	defer GlobalTasksStatuses.Mutex.RUnlock()
	for taskID, processInfo := range GlobalTasksStatuses.Data {
		fmt.Println("key:", taskID, ", val:", processInfo)
		consumer(client, taskID, &processInfo.Result) // don't change last argument
	}
}

func HandleTask(task *pb.TTask) {
	result := pb.TResult{ResultCode: pb.TResult_WAIT.Enum()}
	GlobalTasksStatuses.Store(task.GetId(), ProcessInfo{Result: result})
	executionDir := path.Join(parser.AgentConfig.AgentDirectory, task.GetId())
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

func CreateFile(filePath string) *os.File {
	stdoutFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	return stdoutFile
}

func RunShellCommand(command string, directory string, stdOutFilePath string, stdErrFilePath string, taskId string, changeTaskStatus bool) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = directory
	fmt.Println("Directory ", directory)
	fmt.Println("Command ", command)
	stdoutFile := CreateFile(stdOutFilePath)
	defer stdoutFile.Close()
	cmd.Stdout = stdoutFile
	stderrFile := CreateFile(stdErrFilePath)
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
