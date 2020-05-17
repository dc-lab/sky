package grpc_server

import (
	pb "github.com/dc-lab/sky/api/proto/resource_manager"
	"github.com/dc-lab/sky/resource_manager/app"
	"github.com/dc-lab/sky/resource_manager/db"
	"io"
	"log"
)

type Server struct{}

type AgentConnection struct {
	resourceId string
}

func (s Server) Send(srv pb.ResourceManager_SendServer) error {
	log.Println("start new server")
	ctx := srv.Context()

	var connection AgentConnection

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

		switch x := req.Body.(type) {
		case *pb.TFromAgentMessage_Greetings:
			greetings := req.GetGreetings()
			log.Println("Got greeting request")
			token := greetings.GetToken()
			resourceId, err := db.GetResourceIdByToken(token)
			if err != nil {
				switch err.(type) {
				case *app.ResourceNotFound:
					log.Printf("Can't find resource with token %s", token)
				default:
					log.Printf("Error while authorizing resource")
				}
				continue
			}
			connection.resourceId = resourceId
			log.Printf("Successful authorization for resource %s", resourceId)

			log.Println("Going to send TaskRequest")
			taskId := "123"
			shellCommand := "ls -la && sleep 0.5"
			task := pb.TTask{
				Id:                    taskId,
				ExecutionShellCommand: shellCommand,
			}
			taskRequest := pb.TTaskRequest{Task: &task}
			resp := pb.TToAgentMessage{Body: &pb.TToAgentMessage_TaskRequest{TaskRequest: &taskRequest}}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v\n", err)
			}
			log.Println("send new message")
		case *pb.TFromAgentMessage_HardwareResponse:
			hardwareResp := req.GetHardwareResponse()
			freeHardwareData := hardwareResp.FreeHardwareData
			totalHardwareData := hardwareResp.TotalHardwareData
			log.Printf("Got hardware data: %.2f/%.2f, %d/%d, %d/%d\n",
				freeHardwareData.CoresCount, totalHardwareData.CoresCount,
				freeHardwareData.DiskBytes, totalHardwareData.DiskBytes,
				freeHardwareData.MemoryBytes, totalHardwareData.MemoryBytes)
		case *pb.TFromAgentMessage_TaskResponse:
			taskResponse := req.GetTaskResponse()
			log.Printf("Got task response: %s, %s, %s", taskResponse.GetTaskId(), taskResponse.GetResult().GetErrorCode(), taskResponse.GetResult().GetResultCode())
		case nil:
			log.Println("The field is not set. And that's kind'a strange")
		default:
			log.Printf("TFromAgentMessage.Body has unexpected type %T\n", x)
		}
		//resp := pb.TToAgentMessage{Body: &pb.TToAgentMessage_HardwareRequest{}}
		//if err := srv.Send(&resp); err != nil {
		//	log.Printf("send error %v\n", err)
		//}
		//log.Println("send new message")
	}
}
