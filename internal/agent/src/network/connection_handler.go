package network

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	rm "github.com/dc-lab/sky/api/proto"
	common "github.com/dc-lab/sky/internal/agent/src/common"
	hardware "github.com/dc-lab/sky/internal/agent/src/hardware"
	parser "github.com/dc-lab/sky/internal/agent/src/parser"
)

type Client struct {
	config *parser.Config
	files  *FileStorage
}

func NewClient(config *parser.Config) (*Client, error) {
	files, err := NewFileStorage(config)
	if err != nil {
		return nil, err
	}
	return &Client{config, files}, nil
}

func (c *Client) CreateConnection(address string) (rm.ResourceManager_SendClient, context.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	common.DealWithError(err)

	client := rm.NewResourceManagerClient(conn)
	stream, err := client.Send(context.Background())
	common.DealWithError(err)
	ctx := stream.Context()
	common.DealWithError(err)
	return stream, ctx
}

func (c *Client) ReceiveResourceManagerRequest(client rm.ResourceManager_SendClient) {
	for {
		generalResponse, err := client.Recv()
		if err == io.EOF {
			return
		}
		common.DealWithError(err)
		switch response := generalResponse.Body.(type) {
		case *rm.ToAgentMessage_HardwareRequest:
			log.Debugln("Hardware data request")
			go SendHardwareData(client, hardware.GetTotalHardwareData(c.config.AgentDirectory), hardware.GetFreeHardwareData(c.config.AgentDirectory))
		case *rm.ToAgentMessage_TaskRequest:
			log.Debugln("Task request")
			task := response.TaskRequest.GetTask()
			go StartTask(task, c.config)
		case *rm.ToAgentMessage_StageInRequest:
			log.Debugln("Stage in request")
			files := response.StageInRequest.GetFiles()
			taskId := response.StageInRequest.GetTaskId()
			go c.files.StageInFiles(client, taskId, files)
		case *rm.ToAgentMessage_StageOutRequest:
			log.Debugln("Stage out request")
			taskId := response.StageOutRequest.GetTaskId()
			localPath := response.StageOutRequest.GetAgentRelativeLocalPath()
			go c.files.StageOutFiles(client, taskId, localPath)
		case *rm.ToAgentMessage_CancelTaskRequest:
			taskId := response.CancelTaskRequest.GetTaskId()
			go CancelTask(taskId)
		case *rm.ToAgentMessage_DeleteTaskRequest:
			taskId := response.DeleteTaskRequest.GetTaskId()
			go DeleteTask(taskId)
		default:
			log.Debugln("Non type of response")
		}
	}
	// err = stream.CloseSend()
}

func (c *Client) UpdateHealthFile(healthFilePath string) {
	for ; ; time.Sleep(time.Millisecond * 200) {
		tsString := common.CurrentTimestampMillisString()
		err := ioutil.WriteFile(healthFilePath, []byte(tsString), 0644)
		common.DieWithError(err)
	}
}

func (c *Client) SendHealthChecks(client rm.ResourceManager_SendClient) {
	for ; ; time.Sleep(time.Second * 10) {
		SendHardwareData(client, hardware.GetTotalHardwareData(c.config.AgentDirectory), hardware.GetFreeHardwareData(c.config.AgentDirectory))
	}
}

func (c *Client) Run() error {
	stream, ctx := c.CreateConnection(c.config.ResourceManagerAddress)
	success := ResourceRegistration(stream, c.config.Token)
	if !success {
		log.Println("Failed resource registration. Invalid greetings")
		os.Exit(1)
	}
	go c.ReceiveResourceManagerRequest(stream)
	go c.UpdateHealthFile(c.config.HealthFile)
	go c.SendHealthChecks(stream)
	go UpdateTasksInfo(stream)
	<-ctx.Done()

	return nil
}
