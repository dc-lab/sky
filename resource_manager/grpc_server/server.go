package grpc_server

import (
	"context"
	"fmt"
	"github.com/dc-lab/sky/api/proto/common"
	pb "github.com/dc-lab/sky/api/proto/resource_manager"
	"github.com/dc-lab/sky/resource_manager/app"
	"github.com/dc-lab/sky/resource_manager/db"
	"io"
	"log"
	"time"
)

type Server struct{}

func HandleGreetings(greetings *pb.TGreetings) (*pb.TToAgentMessage, string) {
	log.Println("Got greetings request")
	var resultCode common.TResult_TResultCode
	token := greetings.GetToken()
	var resourceId, err = db.GetResourceIdByToken(token)
	if err != nil {
		resultCode = common.TResult_FAILED
		switch err.(type) {
		case *app.ResourceNotFound:
			log.Printf("Can't find resource with token %s\n", token)
		default:
			log.Printf("Error while authorizing resource: %s\n", err)
		}
	} else {
		connectedAgents.AddAgent(resourceId)
		resultCode = common.TResult_SUCCESS
		log.Printf("Agent for resource %s logged successfuly\n", resourceId)
	}

	result := common.TResult{ResultCode: resultCode}
	greetingsValidation := pb.TGreetingsValidation{Result: &result}
	return &pb.TToAgentMessage{Body: &pb.TToAgentMessage_GreetingsValidation{GreetingsValidation: &greetingsValidation}}, resourceId
}

func HandleHardware(resourceId string, hardware *pb.THardwareResponse) {
	log.Println("Got hardware request")
	if resourceId == "" {
		log.Println("Resource was not registered, so skipping")
		return
	}
	free := hardware.GetFreeHardwareData()
	total := hardware.GetTotalHardwareData()
	if err := connectedAgents.AddHardwareData(resourceId, total, free); err != nil {
		log.Printf("Error while handling hardware request: %s\n", err)
	}
}

func HandleTask(resourceId string, task *pb.TTaskResponse) {
	log.Println("Got task request")
	// todo: transfer data
	log.Printf("Got task response: %s, %s, %s", task.GetTaskId(), task.GetResult().GetErrorCode(), task.GetResult().GetResultCode())
}

func HandleStageIn(resourceId string, stageIn *pb.TStageInResponse) {
	log.Println("Got stage in")
	// todo: transfer data
	log.Printf("Got stage in response: %s, %s, %s", stageIn.GetTaskId(), stageIn.GetResult().GetErrorCode(), stageIn.GetResult().GetResultCode())
}

func HandleStageOut(resourceId string, stageOut *pb.TStageOutResponse) {
	log.Println("Got stage in")
	// todo: transfer data
	log.Printf("Got stage in response: %s", stageOut.GetTaskId())
}

func Healthcheck(resourceId string) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case t := <- ticker.C:
			lastUpdate := connectedAgents.GetLastUpdate(resourceId)
			if lastUpdate == nil {
				log.Printf("Stop checking %s\n", resourceId)
				return
			}
			if time.Since(*lastUpdate).Seconds() > 10 {
				log.Printf("Last update from %s was more than 10 seconds ago, so closing connection", resourceId)
				connectedAgents.RemoveAgent(resourceId)
				return
			}
			log.Printf("ResourceId: %s, tick at %s", resourceId, t)
		}
	}
}

func (s Server) Send(srv pb.ResourceManager_SendServer) error {
	log.Println("start new server")
	ctx := srv.Context()

	var resourceId string

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v\n", err)
			continue
		}

		var response *pb.TToAgentMessage = nil

		switch x := req.Body.(type) {
		case *pb.TFromAgentMessage_Greetings:
			response, resourceId = HandleGreetings(req.GetGreetings())
			if resourceId != "" {
				go Healthcheck(resourceId)
			}
		case *pb.TFromAgentMessage_HardwareResponse:
			HandleHardware(resourceId, req.GetHardwareResponse())
		case *pb.TFromAgentMessage_TaskResponse:
			HandleTask(resourceId, req.GetTaskResponse())
		case *pb.TFromAgentMessage_StageInResponse:
			HandleStageIn(resourceId, req.GetStageInResponse())
		case *pb.TFromAgentMessage_StageOutResponse:
			HandleStageOut(resourceId, req.GetStageOutResponse())
		default:
			log.Printf("TFromAgentMessage.Body has unexpected type %T\n", x)
		}

		if response != nil {
			if err := srv.Send(response); err != nil {
				log.Printf("Error while sending message to agent: %v\n", err)
			}
		}

		if resourceId != "" {
			message := connectedAgents.GetMessage(resourceId)
			if message != nil {
				if err := srv.Send(message); err != nil {
					log.Printf("Error while sending message to agent: %v\n", err)
				}
			}
		}
	}
}

func (s Server) Update(ctx context.Context, request *pb.TResourceRequest) (*pb.TResourceResponse, error) {
	panic("implement me")
}

func (s Server) AgentAction(ctx context.Context, request *pb.TRMRequest) (*pb.TRMResponse, error) {
	resourceId := request.GetResourceId()
	body := request.GetRealMessage()
	err := connectedAgents.AddMessage(resourceId, body)
	if err != nil {
		log.Printf("Error during adding message to resoure %s: %v\n", resourceId, err)
		return &pb.TRMResponse{ResultCode: pb.TRMResponse_NOT_FOUND}, err
	}
	return &pb.TRMResponse{ResultCode: pb.TRMResponse_OK}, err
}
