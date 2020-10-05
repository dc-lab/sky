package job_manager

import (
	"context"
	"errors"
	"fmt"
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

func (d *dispatcher) Run() {
	err := d.run()
	log.WithError(err).Fatalln("dispatcher failed")
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
	agentMessage := &pb.TToAgentMessage_TaskRequest{}

	res, err := d.rmClient.AgentAction(context.Background(), &pb.TRMRequest{
		ResourceId: task.resource,
		RealMessage: &pb.TToAgentMessage{
			Body: agentMessage,
		},
	})

	if err != nil {
		log.WithError(err).Errorf("Start task request failed")
		return err
	}

	if code := res.GetResultCode(); code != pb.TRMResponse_OK {
		log.WithField("code", code.String()).Errorf("Start task request failed: bad return code")
		return fmt.Errorf("Start task request failed")
	}

	return nil
}

func (d *dispatcher) PushTask(task assignedTask) error {
	d.startedTasks <- task
	return nil
}
