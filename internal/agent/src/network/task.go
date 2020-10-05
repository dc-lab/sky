package network

import (
	pb "github.com/dc-lab/sky/api/proto"
	rm "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/agent/src/common"
	"github.com/dc-lab/sky/internal/agent/src/executors"
	"github.com/dc-lab/sky/internal/agent/src/parser"
	"path"
	"sync/atomic"
	"time"
)

type Task struct {
	TaskId       string                 // don't change
	ExecutionDir string                 // don't change
	Executor     executors.TaskExecutor // don't change
	Result       *pb.TResult
	IsFinished   atomic.Value
	QuitChanel   chan struct{}
}

func GetTaskExecutionDir(taskId, dir string) (string, error) {
	executionDir := path.Join(dir, taskId)
	err := error(nil)
	if val, err := common.PathExists(executionDir, true); !val && err == nil {
		err = common.CreateDirectory(executionDir, false)
	}
	return executionDir, err
}

func (t *Task) Init(taskProto *rm.TTask, config *parser.Config) {
	t.QuitChanel = make(chan struct{}, 1)
	t.IsFinished.Store(false)
	t.TaskId = taskProto.GetId()
	var err error
	t.ExecutionDir, err = GetTaskExecutionDir(taskProto.GetId(), config.AgentDirectory)
	common.DealWithError(err)
	if taskProto.GetDockerImage() != "" {
		t.Executor = &executors.DockerExecutor{
			Image:                    taskProto.GetDockerImage(),
			RequirementsShellCommand: taskProto.GetRequirementsShellCommand(),
			ExecutionShellCommand:    taskProto.GetExecutionShellCommand(),
			ExecutionDir:             t.ExecutionDir,
		}
	} else {
		t.Executor = &executors.LocalExecutor{
			RequirementsShellCommand: taskProto.GetRequirementsShellCommand(),
			ExecutionShellCommand:    taskProto.GetExecutionShellCommand(),
			ExecutionDir:             t.ExecutionDir,
		}
	}
}

func (t *Task) InstallRequirements() {
	updateFinalResultFunc := func(err error) {
		if err != nil {
			result := pb.TResult{ResultCode: pb.TResult_FAILED, ErrorCode: pb.TResult_INTERNAL}
			common.DealWithError(err)
			GlobalTasksStatuses.SetTaskResult(t.TaskId, &result)
		}
	}
	t.Executor.Prepare(updateFinalResultFunc)
}

func (t *Task) Run() {
	updateFinalResultFunc := func(err error) {
		result := pb.TResult{ResultCode: pb.TResult_SUCCESS}
		if err != nil {
			result = pb.TResult{ResultCode: pb.TResult_FAILED, ErrorCode: pb.TResult_INTERNAL}
			common.DealWithError(err)
		}
		GlobalTasksStatuses.SetTaskResult(t.TaskId, &result)
		t.IsFinished.Store(true)
	}
	getProcessInfoBeforeExecution := func(result *pb.TResult) {
		GlobalTasksStatuses.SetTaskResult(t.TaskId, result)
	}
	t.Executor.Run(t.QuitChanel, getProcessInfoBeforeExecution, updateFinalResultFunc)
}

func (t *Task) Cancel() {
	select {
	case <-time.After(2 * time.Second):
		return
	case t.QuitChanel <- struct{}{}:
		result := pb.TResult{ResultCode: pb.TResult_CANCELED}
		GlobalTasksStatuses.SetTaskResult(t.TaskId, &result)
		t.IsFinished.Store(true)
	}
}

func (t *Task) Delete() {
	t.Cancel()
	err := common.RemoveDirectory(t.ExecutionDir)
	common.DealWithError(err)
	result := pb.TResult{ResultCode: pb.TResult_DELETED}
	GlobalTasksStatuses.SetTaskResult(t.TaskId, &result)
}
