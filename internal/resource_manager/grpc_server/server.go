package grpc_server

import (
	"context"
	"io"
	"time"

	dm "github.com/dc-lab/sky/api/proto"
	pb "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/resource_manager/app"
	"github.com/dc-lab/sky/internal/resource_manager/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Server struct {
	// FIXME: We should not push rm -> dm
	// just cache agents responses on rm and pull them from dm
	dmClient dm.DataManagerClient
}

func NewServer(dmAddress string) (*Server, error) {
	conn, err := grpc.Dial(dmAddress, grpc.WithInsecure(), grpc.WithBackoffMaxDelay(time.Second*10))
	if err != nil {
		return nil, err
	}
	client := dm.NewDataManagerClient(conn)
	return &Server{client}, nil
}

func HandleGreetings(greetings *pb.Greetings) (*pb.ToAgentMessage, string) {
	log.Println("Got greetings request")
	var resultCode pb.Result_ResultCode
	token := greetings.GetToken()
	var resourceId, err = db.GetResourceIdByToken(token)
	if err != nil {
		resultCode = pb.Result_FAILED
		switch err.(type) {
		case *app.ResourceNotFound:
			log.Printf("Can't find resource with token %s\n", token)
		default:
			log.Printf("Error while authorizing resource: %s\n", err)
		}
	} else {
		db.ConnectedAgents.AddAgent(resourceId)
		resultCode = pb.Result_SUCCESS
		log.Printf("Agent for resource %s logged successfuly\n", resourceId)
	}

	result := pb.Result{ResultCode: resultCode}
	greetingsValidation := pb.GreetingsValidation{Result: &result}
	return &pb.ToAgentMessage{Body: &pb.ToAgentMessage_GreetingsValidation{GreetingsValidation: &greetingsValidation}}, resourceId
}

func HandleHardware(resourceId string, hardware *pb.HardwareResponse) {
	log.Println("Got hardware request")
	if resourceId == "" {
		log.Println("Resource was not registered, so skipping")
		return
	}
	free := hardware.GetFreeHardwareData()
	total := hardware.GetTotalHardwareData()
	if err := db.ConnectedAgents.AddHardwareData(resourceId, total, free); err != nil {
		log.Printf("Error while handling hardware request: %s\n", err)
	}
}

func HandleTask(resourceId string, task *pb.TaskResponse, dmClient dm.DataManagerClient) {
	log.Println("Got task request")
	result := &dm.UpdateTaskResultsRequest{
		Files:   task.TaskFiles,
		TaskId:  task.TaskId,
		AgentId: resourceId,
	}
	_, err := dmClient.UpdateTaskResults(context.Background(), result)
	if err != nil {
		log.Printf("UpdateTaskResults request failed: %e", err)
	}
	log.Printf("Got task response: %s, %s, %s", task.GetTaskId(), task.GetResult().GetErrorCode(), task.GetResult().GetResultCode())
}

func HandleStageIn(resourceId string, stageIn *pb.StageInResponse) {
	log.Println("Got stage in")
	// todo: transfer data
	log.Printf("Got stage in response: %s, %s, %s", stageIn.GetTaskId(), stageIn.GetResult().GetErrorCode(), stageIn.GetResult().GetResultCode())
}

func HandleStageOut(resourceId string, stageOut *pb.StageOutResponse) {
	log.Println("Got stage in")
	// todo: transfer data
	log.Printf("Got stage in response: %s", stageOut.GetTaskId())
}

func Healthcheck(resourceId string) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case t := <-ticker.C:
			lastUpdate := db.ConnectedAgents.GetLastUpdate(resourceId)
			if lastUpdate == nil {
				log.Printf("Stop checking %s\n", resourceId)
				return
			}
			if time.Since(*lastUpdate).Seconds() > 10 {
				log.Printf("Last update from %s was more than 10 seconds ago, so closing connection", resourceId)
				db.ConnectedAgents.RemoveAgent(resourceId)
				return
			}
			log.Printf("ResourceId: %s, tick at %s", resourceId, t)
		}
	}
}

func (s *Server) Send(srv pb.ResourceManager_SendServer) error {
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

		var response *pb.ToAgentMessage = nil

		switch x := req.Body.(type) {
		case *pb.FromAgentMessage_Greetings:
			response, resourceId = HandleGreetings(req.GetGreetings())
			if resourceId != "" {
				go Healthcheck(resourceId)
			}
		case *pb.FromAgentMessage_HardwareResponse:
			HandleHardware(resourceId, req.GetHardwareResponse())
		case *pb.FromAgentMessage_TaskResponse:
			HandleTask(resourceId, req.GetTaskResponse(), s.dmClient)
		case *pb.FromAgentMessage_StageInResponse:
			HandleStageIn(resourceId, req.GetStageInResponse())
		case *pb.FromAgentMessage_StageOutResponse:
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
			message := db.ConnectedAgents.GetMessage(resourceId)
			if message != nil {
				if err := srv.Send(message); err != nil {
					log.Printf("Error while sending message to agent: %v\n", err)
				}
			}
		}
	}
}

func (s *Server) Update(ctx context.Context, request *pb.ResourceRequest) (*pb.ResourceResponse, error) {
	switch x := request.Body.(type) {
	case *pb.ResourceRequest_CreateResourceRequest:
		pbResource := request.GetCreateResourceRequest().GetResource()
		resourceId := pbResource.Id
		resource := db.Resource{
			Id:    resourceId,
			Name:  pbResource.Id,
			Type:  db.GetStringTypeByEnum(pbResource.Type),
			Owner: pbResource.OwnerId,
			Token: pbResource.Token,
		}
		err := resource.CreateMe()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		err = db.AddUsersToResource(resourceId, pbResource.Permissions.GetUsers())
		if err != nil {
			log.Println(err)
			return nil, err
		}
		err = db.AddGroupsToResource(resourceId, pbResource.Permissions.GetGroups())
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &pb.ResourceResponse{Body: &pb.ResourceResponse_CreateResourceResponse{}}, nil
	case *pb.ResourceRequest_DeleteResourceRequest:
		resourceId := request.GetDeleteResourceRequest().ResourceId
		userId := request.GetDeleteResourceRequest().UserId
		if err := db.DeleteResource(userId, resourceId); err != nil {
			log.Println(err)
			return nil, err
		}
		return &pb.ResourceResponse{Body: &pb.ResourceResponse_DeleteResourceResponse{}}, nil
	default:
		log.Printf("TResourceRequest.Body has unexpected type %T\n", x)
		return nil, nil
	}
}

func (s *Server) AgentAction(ctx context.Context, request *pb.RMRequest) (*pb.RMResponse, error) {
	resourceId := request.GetResourceId()
	body := request.GetRealMessage()
	err := db.ConnectedAgents.AddMessage(resourceId, body)
	if err != nil {
		log.Printf("Error during adding message to resoure %s: %v\n", resourceId, err)
		return &pb.RMResponse{ResultCode: pb.RMResponse_NOT_FOUND}, err
	}
	return &pb.RMResponse{ResultCode: pb.RMResponse_OK}, err
}

func (s *Server) GetResources(ctx context.Context, request *pb.GetResourcesRequest) (*pb.GetResourcesResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "method is not implemented")
}
