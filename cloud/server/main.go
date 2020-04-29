package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"

	pb "github.com/dc-lab/sky/cloud/proto"
	"github.com/dc-lab/sky/cloud/server/cmd"
	"github.com/dc-lab/sky/cloud/server/handlers"
)

type Server struct{}

func (s Server) DoAction(srv pb.TCloudConnector_DoActionServer) error {
	log.Println("start new server")
	ctx := srv.Context()

	for {
		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}

		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		// parse request type from received stream
		var resp *pb.TCloudResponse
		switch typedReq := req.Body.(type) {
		case *pb.TCloudRequest_AllocateInstanceRequest:
			typedResp := handlers.HandleAllocateVMRequest(typedReq.AllocateInstanceRequest)
			resp = &pb.TCloudResponse{
				Body: &pb.TCloudResponse_AllocateInstanceResponse{AllocateInstanceResponse: &typedResp},
			}
		case *pb.TCloudRequest_DeallocateInstanceRequest:
			typedResp := handlers.HandleDeallocateVMRequest(typedReq.DeallocateInstanceRequest)
			resp = &pb.TCloudResponse{
				Body: &pb.TCloudResponse_DeallocateInstanceResponse{DeallocateInstanceResponse: &typedResp},
			}
		case *pb.TCloudRequest_ConnectInstanceRequest:
			typedResp := handlers.HandleConnectInstanceRequest(typedReq.ConnectInstanceRequest)
			resp = &pb.TCloudResponse{
				Body: &pb.TCloudResponse_ConnectInstanceResponse{ConnectInstanceResponse: &typedResp},
			}
		case *pb.TCloudRequest_DisconnectInstanceRequest:
			typedResp := handlers.HandleDisconnectInstanceRequest(typedReq.DisconnectInstanceRequest)
			resp = &pb.TCloudResponse{
				Body: &pb.TCloudResponse_DisconnectInstanceResponse{DisconnectInstanceResponse: &typedResp},
			}
		case *pb.TCloudRequest_StartInstanceRequest:
			typedResp := handlers.HandleStartVMRequest(typedReq.StartInstanceRequest)
			resp = &pb.TCloudResponse{
				Body: &pb.TCloudResponse_StartInstanceResponse{StartInstanceResponse: &typedResp},
			}
		case *pb.TCloudRequest_StopInstanceRequest:
			typedResp := handlers.HandleStopVMRequest(typedReq.StopInstanceRequest)
			resp = &pb.TCloudResponse{
				Body: &pb.TCloudResponse_StopInstanceResponse{StopInstanceResponse: &typedResp},
			}
		case *pb.TCloudRequest_DeployImageRequest:
			typedResp := handlers.HandleDeployImageRequest(typedReq.DeployImageRequest)
			resp = &pb.TCloudResponse{
				Body: &pb.TCloudResponse_DeployImageResponse{DeployImageResponse: &typedResp},
			}
		default:
			log.Print("got unrecognized req")
		}

		// send it to stream
		if err := srv.Send(resp); err != nil {
			log.Printf("send error %v", err)
		}
	}
}

// Usage:
// go run server.go [--grpc-port PORT]
func main() {
	cmd.Execute()

	// create listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cmd.GrpcPort))
	log.Printf("listen tcp on %d", cmd.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterTCloudConnectorServer(s, Server{})

	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
