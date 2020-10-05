package job_manager

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	pb "github.com/dc-lab/sky/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type assignedTask struct {
	resource string
	task     *Task
}

type dispatcher struct {
	rmClient     pb.ResourceManagerClient
	startedTasks chan assignedTask
}

func NewDispatcher(config *Config) (*dispatcher, error) {
	conn, err := grpc.Dial(config.ResourceManagerAddress, grpc.WithInsecure(), grpc.WithBackoffMaxDelay(time.Second*10))
	if err != nil {
		return nil, err
	}
	client := pb.NewResourceManagerClient(conn)
	return &dispatcher{
		rmClient:     client,
		startedTasks: make(chan assignedTask),
	}, nil
}

func (d *dispatcher) Run(wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()
	err := d.run()
	log.WithError(err).Errorf("dispatcher failed")
	return err
}

func (d *dispatcher) run() error {
	for {
		select {
		case task := <-d.startedTasks:
			d.startTask(task)
		}
	}
}

func (d *dispatcher) startTask(task assignedTask) error {
	err := d.stageInTaskFiles(task)
	if err != nil {
		return err
	}

	err = d.sendTaskToAgent(task)
	if err != nil {
		return err
	}

	return nil
}

func (d *dispatcher) agentAction(agent string, message *pb.ToAgentMessage) error {
	res, err := d.rmClient.AgentAction(context.Background(), &pb.RMRequest{
		ResourceId:  agent,
		RealMessage: message,
	})

	if err != nil {
		log.WithError(err).Errorf("Agent request failed")
		return err
	}

	if code := res.GetResultCode(); code != pb.RMResponse_OK {
		log.WithField("code", code.String()).Errorf("Agent request failed: bad return code")
		return fmt.Errorf("Agent request failed")
	}

	return nil
}

func (d *dispatcher) stageInTaskFiles(task assignedTask) error {
	files := make([]*pb.FileOnAgent, len(task.task.Files))
	for i := range task.task.Files {
		files[i] = &pb.FileOnAgent{
			Id:                     task.task.Files[i],
			AgentRelativeLocalPath: task.task.FilePaths[i],
		}
	}

	agentMessage := &pb.ToAgentMessage_StageInRequest{
		StageInRequest: &pb.StageInRequest{
			TaskId: task.task.ID.String(),
			Files:  files,
		},
	}

	err := d.agentAction(task.resource, &pb.ToAgentMessage{Body: agentMessage})
	if err != nil {
		return err
	}

	return nil
}

func (d *dispatcher) sendTaskToAgent(task assignedTask) error {
	shellCommandBuilder := strings.Builder{}
	for i, arg := range task.task.Command {
		if i > 0 {
			shellCommandBuilder.WriteString(" ")
		}
		shellCommandBuilder.WriteString(arg)
	}

	agentMessage := &pb.ToAgentMessage_TaskRequest{
		TaskRequest: &pb.TaskRequest{
			Task: &pb.Task{
				Id:                    task.task.ID.String(),
				ExecutionShellCommand: shellCommandBuilder.String(),
			},
		},
	}

	err := d.agentAction(task.resource, &pb.ToAgentMessage{Body: agentMessage})
	if err != nil {
		return err
	}

	return nil
}

func (d *dispatcher) PushTask(task assignedTask) error {
	d.startedTasks <- task
	return nil
}
