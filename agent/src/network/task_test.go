package network

import (
	"github.com/dc-lab/sky/agent/src/common"
	"github.com/dc-lab/sky/agent/src/executors"
	"github.com/dc-lab/sky/agent/src/parser"
	pb "github.com/dc-lab/sky/api/proto/common"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
	"github.com/docker/distribution/uuid"
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func InitializeAgentConfig() {
	parser.InitializeAgentConfigFromCustomFields(map[string]interface{}{
		"LogsDirectory":          "/tmp/agent-logs-test",
		"RunDirectory":           "/tmp/agent_test/",
		"ResourceManagerAddress": "localhost:5051",
		"AgentDirectory":         "/tmp/agent_test",
		"TokenPath":              "/tmp/sample_token",
	})
}

func TestGetTaskExecutionDir(t *testing.T) {
	InitializeAgentConfig()
	taskId := uuid.Generate().String()
	executionDir, err := GetTaskExecutionDir(taskId)
	assert.Equal(t, err, nil)
	existFlag, err := common.PathExists(executionDir, true)
	assert.Equal(t, err, nil)
	assert.Equal(t, existFlag, true)
}

func GetLocalTaskProto(taskId string) rm.TTask {
	requirementsCommand := "apt-get install python"
	executionShellCommand := "python -c 'print(5)'"
	taskProto := rm.TTask{
		Id:                       taskId,
		RequirementsShellCommand: requirementsCommand,
		ExecutionShellCommand:    executionShellCommand,
	}
	return taskProto
}

func GetDockerTask(taskId string) rm.TTask {
	taskProto := GetLocalTaskProto(taskId)
	taskProto.DockerImage = "image"
	return taskProto
}

func TestTask_Init(t *testing.T) {
	InitializeAgentConfig()
	taskId := uuid.Generate().String()
	var task Task
	taskProto := GetDockerTask(taskId)
	task.Init(&taskProto)
	GlobalTasksStatuses.Store(task.TaskId, &task)
	assert.Equal(t, task.IsFinished.Load(), false)
	assert.Equal(t, task.TaskId, taskId)
	executionDir, err := GetTaskExecutionDir(taskId)
	assert.Equal(t, err, nil)
	executor := task.Executor.(*executors.DockerExecutor)
	assert.Equal(t, executor.ExecutionDir, executionDir)
	assert.Equal(t, executor.RequirementsShellCommand, taskProto.GetRequirementsShellCommand())
	assert.Equal(t, executor.ExecutionShellCommand, taskProto.GetExecutionShellCommand())
	assert.Equal(t, executor.Image, taskProto.GetDockerImage())
}

func TestTask_InstallRequirements(t *testing.T) {
	InitializeAgentConfig()
	taskId := uuid.Generate().String()
	result := pb.TResult{ResultCode: pb.TResult_WAIT}
	task := Task{Result: &result}
	taskProto := GetLocalTaskProto(taskId)
	task.Init(&taskProto)
	GlobalTasksStatuses.Store(task.TaskId, &task)
	executor := task.Executor.(*executors.LocalExecutor)
	pid := executor.ProcessID
	task.InstallRequirements()
	assert.NotEqual(t, executor.ProcessID, pid)
	assert.FileExists(t, path.Join(task.ExecutionDir, "requirements_out"))
	assert.FileExists(t, path.Join(task.ExecutionDir, "requirements_err"))
}
