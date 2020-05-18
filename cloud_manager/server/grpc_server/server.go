package grpc_server

import (
	"io"
	"log"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

type Server struct{}

func (s Server) DoAction(srv cm.TCloudManager_DoActionServer) error {
	log.Println("Start new gRPC connection")
	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err == io.EOF {
			// close stream from server side
			log.Println("Connection closed. Exiting...")
			return nil
		}

		if err != nil {
			log.Printf("Receive error %v", err)
			continue
		}

		// parse request type from received stream
		var resp *cm.TCloudResponse
		switch typedReq := req.Body.(type) {
		case *cm.TCloudRequest_AllocateInstanceRequest:
			typedResp := HandleAllocateVMRequest(typedReq.AllocateInstanceRequest)
			resp = &cm.TCloudResponse{
				Body: &cm.TCloudResponse_AllocateInstanceResponse{AllocateInstanceResponse: &typedResp},
			}
		case *cm.TCloudRequest_DeallocateInstanceRequest:
			typedResp := HandleDeallocateVMRequest(typedReq.DeallocateInstanceRequest)
			resp = &cm.TCloudResponse{
				Body: &cm.TCloudResponse_DeallocateInstanceResponse{DeallocateInstanceResponse: &typedResp},
			}
		case *cm.TCloudRequest_ConnectInstanceRequest:
			typedResp := HandleConnectInstanceRequest(typedReq.ConnectInstanceRequest)
			resp = &cm.TCloudResponse{
				Body: &cm.TCloudResponse_ConnectInstanceResponse{ConnectInstanceResponse: &typedResp},
			}
		case *cm.TCloudRequest_DisconnectInstanceRequest:
			typedResp := HandleDisconnectInstanceRequest(typedReq.DisconnectInstanceRequest)
			resp = &cm.TCloudResponse{
				Body: &cm.TCloudResponse_DisconnectInstanceResponse{DisconnectInstanceResponse: &typedResp},
			}
		case *cm.TCloudRequest_StartInstanceRequest:
			typedResp := HandleStartVMRequest(typedReq.StartInstanceRequest)
			resp = &cm.TCloudResponse{
				Body: &cm.TCloudResponse_StartInstanceResponse{StartInstanceResponse: &typedResp},
			}
		case *cm.TCloudRequest_StopInstanceRequest:
			typedResp := HandleStopVMRequest(typedReq.StopInstanceRequest)
			resp = &cm.TCloudResponse{
				Body: &cm.TCloudResponse_StopInstanceResponse{StopInstanceResponse: &typedResp},
			}
		case *cm.TCloudRequest_DeployImageRequest:
			typedResp := HandleDeployImageRequest(typedReq.DeployImageRequest)
			resp = &cm.TCloudResponse{
				Body: &cm.TCloudResponse_DeployImageResponse{DeployImageResponse: &typedResp},
			}
		default:
			log.Printf("Got unrecognized request: %v", req.Body)
		}

		// send it to stream
		if err := srv.Send(resp); err != nil {
			log.Printf("Send error %v", err)
		}
	}
}
