package grpc_server

import (
	pb "github.com/dc-lab/sky/api/proto/resource_manager"
	"io"
	"log"
	"strconv"
)

type Server struct{}

func (s Server) Send(srv pb.ResourceManager_SendServer) error {
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

		var taskId int8
		switch x := req.Body.(type) {
		case *pb.TFromAgentMessage_Greetings:
			greetings := req.GetGreetings()
			log.Printf("Got greetings: %s", greetings.GetToken())
			log.Println("Going to send TaskRequest")
			taskId++
			shellCommand := "ls -la && sleep 0.5"
			task := pb.TTask{
				Id:                       strconv.Itoa(int(taskId)),
				ExecutionShellCommand:    shellCommand,
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
