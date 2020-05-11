package network

import (
	"github.com/dc-lab/sky/agent/src/common"
	"github.com/dc-lab/sky/agent/src/parser"
	pb "github.com/dc-lab/sky/api/proto/common"
	"os"
	"path"
)

type Task struct {
	TaskId       string
	ExecutionDir string
	Result       *pb.TResult
	ProcessID    int
}

func (t *Task) Init() {
	t.ExecutionDir = path.Join(parser.AgentConfig.AgentDirectory, t.TaskId)
	err := error(nil)
	if val, err := common.PathExists(t.ExecutionDir, true); !val && err == nil {
		err = os.Mkdir(t.ExecutionDir, 0777)
	}
	common.DealWithError(err)
}
