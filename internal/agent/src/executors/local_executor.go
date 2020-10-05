package executors

import (
	"github.com/dc-lab/sky/internal/agent/src/common"
	pb "github.com/dc-lab/sky/api/proto"
	"os/exec"
	"path"
	"sync/atomic"
)

type TaskExecutor interface {
	Prepare(func(err error))
	Run(<-chan struct{}, func(result *pb.TResult), func(err error))
}

type LocalExecutor struct {
	ExecutionDir             string // don't change
	RequirementsShellCommand string // don't change
	ExecutionShellCommand    string // don't change
	ProcessID                int64  // use atomic
}

func (e *LocalExecutor) Prepare(afterExecution func(err error)) {
	if e.RequirementsShellCommand != "" {
		e.RunShellCommand(
			e.RequirementsShellCommand,
			e.ExecutionDir,
			path.Join(e.ExecutionDir, "requirements_out"),
			path.Join(e.ExecutionDir, "requirements_err"),
			nil,
			afterExecution,
			nil)
	}
}

func (e *LocalExecutor) Run(
	quiteChannel <-chan struct{},
	beforeExecution func(result *pb.TResult),
	afterExecution func(err error),
) {
	e.RunShellCommand(
		e.ExecutionShellCommand,
		e.ExecutionDir,
		path.Join(e.ExecutionDir, "execution_out"),
		path.Join(e.ExecutionDir, "execution_err"),
		beforeExecution,
		afterExecution,
		quiteChannel,
	)
}

func (e *LocalExecutor) RunShellCommand(
	command string,
	directory string,
	stdOutFilePath string,
	stdErrFilePath string,
	beforeExecution func(result *pb.TResult),
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
	atomic.StoreInt64(&e.ProcessID, int64(cmd.Process.Pid))
	result := pb.TResult{ResultCode: pb.TResult_RUN}
	if beforeExecution != nil {
		beforeExecution(&result)
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
