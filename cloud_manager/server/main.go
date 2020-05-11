package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"

	cloud "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/cloud_manager/server/cmd"
	"github.com/dc-lab/sky/cloud_manager/server/handlers"
)

type Server struct{}

func (s Server) DoAction(srv cloud.TCloudManager_DoActionServer) error {
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
		var resp *cloud.TCloudResponse
		switch typedReq := req.Body.(type) {
		case *cloud.TCloudRequest_AllocateInstanceRequest:
			typedResp := handlers.HandleAllocateVMRequest(typedReq.AllocateInstanceRequest)
			resp = &cloud.TCloudResponse{
				Body: &cloud.TCloudResponse_AllocateInstanceResponse{AllocateInstanceResponse: &typedResp},
			}
		case *cloud.TCloudRequest_DeallocateInstanceRequest:
			typedResp := handlers.HandleDeallocateVMRequest(typedReq.DeallocateInstanceRequest)
			resp = &cloud.TCloudResponse{
				Body: &cloud.TCloudResponse_DeallocateInstanceResponse{DeallocateInstanceResponse: &typedResp},
			}
		case *cloud.TCloudRequest_ConnectInstanceRequest:
			typedResp := handlers.HandleConnectInstanceRequest(typedReq.ConnectInstanceRequest)
			resp = &cloud.TCloudResponse{
				Body: &cloud.TCloudResponse_ConnectInstanceResponse{ConnectInstanceResponse: &typedResp},
			}
		case *cloud.TCloudRequest_DisconnectInstanceRequest:
			typedResp := handlers.HandleDisconnectInstanceRequest(typedReq.DisconnectInstanceRequest)
			resp = &cloud.TCloudResponse{
				Body: &cloud.TCloudResponse_DisconnectInstanceResponse{DisconnectInstanceResponse: &typedResp},
			}
		case *cloud.TCloudRequest_StartInstanceRequest:
			typedResp := handlers.HandleStartVMRequest(typedReq.StartInstanceRequest)
			resp = &cloud.TCloudResponse{
				Body: &cloud.TCloudResponse_StartInstanceResponse{StartInstanceResponse: &typedResp},
			}
		case *cloud.TCloudRequest_StopInstanceRequest:
			typedResp := handlers.HandleStopVMRequest(typedReq.StopInstanceRequest)
			resp = &cloud.TCloudResponse{
				Body: &cloud.TCloudResponse_StopInstanceResponse{StopInstanceResponse: &typedResp},
			}
		case *cloud.TCloudRequest_DeployImageRequest:
			typedResp := handlers.HandleDeployImageRequest(typedReq.DeployImageRequest)
			resp = &cloud.TCloudResponse{
				Body: &cloud.TCloudResponse_DeployImageResponse{DeployImageResponse: &typedResp},
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
	cloud.RegisterTCloudManagerServer(s, Server{})

	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
