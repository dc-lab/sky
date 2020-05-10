package main

import (
	pb "github.com/dc-lab/sky/resource_manager_grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct{}

func (s server) Send(srv pb.ResourceManager_SendServer) error {
	log.Println("start new server")
	ctx := srv.Context()

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
			log.Printf("Got greetings: %s", greetings.GetToken())
			log.Println("Going to send TaskRequest")
			taskId := "123"
			shellCommand := []byte("ls -la")
			task := pb.TTask{
				Id:                       &taskId,
				ExecutionShellCommand:    shellCommand,
				RequirementsShellCommand: nil,
			}
			taskRequest := pb.TTaskRequest{Task: &task}
			resp := pb.TToAgentMessage{Body: &pb.TToAgentMessage_TaskRequest{TaskRequest: &taskRequest}}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v\n", err)
			}
			log.Println("send new message")
		case *pb.TFromAgentMessage_HardwareData:
			hardwareData := req.GetHardwareData()
			log.Printf("Got hardware data: %d, %d, %d\n", hardwareData.GetCoresCount(), hardwareData.GetDiskAmount(), hardwareData.GetMemoryAmount())
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

func main() {
	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterResourceManagerServer(s, server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
