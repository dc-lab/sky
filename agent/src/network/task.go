package network

import (
	"github.com/dc-lab/sky/agent/src/common"
	"github.com/dc-lab/sky/agent/src/parser"
	pb "github.com/dc-lab/sky/api/proto/common"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
	"os"
	"path"
	"sync/atomic"
)

type Task struct {
	TaskId                   string // don't change
	ExecutionDir             string // don't change
	RequirementsShellCommand string // don't change
	ExecutionShellCommand    string // don't change
	Result                   *pb.TResult
	ProcessID                int64 // use atomic
	IsFinished               atomic.Value
}

func (t *Task) Init(taskProto *rm.TTask) {
	t.IsFinished.Store(false)
	t.TaskId = taskProto.GetId()
	t.ExecutionShellCommand = taskProto.GetExecutionShellCommand()
	t.RequirementsShellCommand = taskProto.GetRequirementsShellCommand()
	t.ExecutionDir = path.Join(parser.AgentConfig.AgentDirectory, t.TaskId)
	err := error(nil)
	if val, err := common.PathExists(t.ExecutionDir, true); !val && err == nil {
		err = os.Mkdir(t.ExecutionDir, 0777)
	}
	common.DealWithError(err)
}

func (t *Task) InstallRequirements() {
	RunShellCommand(
		t.RequirementsShellCommand,
		t.ExecutionDir,
		path.Join(t.ExecutionDir, "requirements_out"),
		path.Join(t.ExecutionDir, "requirements_err"),
		nil,
		nil)
}

func (t *Task) Run() {
	updateFinalResultFunc := func(err error) {
		result := pb.TResult{ResultCode: pb.TResult_SUCCESS}
		if err != nil {
			result = pb.TResult{ResultCode: pb.TResult_FAILED, ErrorCode: pb.TResult_INTERNAL}
			common.DealWithError(err)
		}
		GlobalTasksStatuses.UpdateTaskResult(t.TaskId, &result)
		t.IsFinished.Store(true)
	}
	getProcessInfoBeforeExecution := func(pid int64, result *pb.TResult) {
		GlobalTasksStatuses.UpdateTaskResult(t.TaskId, result)
		atomic.StoreInt64(&t.ProcessID, pid)
	}
	RunShellCommand(
		t.ExecutionShellCommand,
		t.ExecutionDir,
		path.Join(t.ExecutionDir, "execution_out"),
		path.Join(t.ExecutionDir, "execution_err"),
		getProcessInfoBeforeExecution,
		updateFinalResultFunc)
}

func (t *Task) Cancel() {

}
