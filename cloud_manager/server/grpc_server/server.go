package grpc_server

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/cloud_manager/server/grpc_server/async"
	"github.com/dc-lab/sky/cloud_manager/server/grpc_server/ping"
)

type Server struct{}

func (s Server) DoAsync(_ context.Context, req *cm.TAsyncCloudRequest) (*cm.TAsyncCloudResponse, error) {
	log.Println("Got async cloud request")
	switch typedReq := req.Body.(type) {
	case *cm.TAsyncCloudRequest_AllocateInstanceRequest:
		return async.HandleAllocateInstanceRequest(typedReq.AllocateInstanceRequest)
	case *cm.TAsyncCloudRequest_DeallocateInstanceRequest:
		return async.HandleDeallocateInstanceRequest(typedReq.DeallocateInstanceRequest)
	case *cm.TAsyncCloudRequest_StartInstanceRequest:
		return async.HandleStartInstanceRequest(typedReq.StartInstanceRequest)
	case *cm.TAsyncCloudRequest_StopInstanceRequest:
		return async.HandleStopInstanceRequest(typedReq.StopInstanceRequest)
	case *cm.TAsyncCloudRequest_StartTaskRequest:
		return async.HandleStartTaskRequest(typedReq.StartTaskRequest)
	case *cm.TAsyncCloudRequest_StopTaskRequest:
		return async.HandleStopTaskRequest(typedReq.StopTaskRequest)
	case *cm.TAsyncCloudRequest_DeployImageRequest:
		return async.HandleDeployImageRequest(typedReq.DeployImageRequest)
	default:
		log.Printf("Got unrecognized request: %v", req)
		return nil, status.Error(codes.InvalidArgument, "Got unrecognized request")
	}
}

func (s Server) PingTransaction(_ context.Context, req *cm.TPingTransactionRequest) (*cm.TPingTransactionResponse, error) {
	log.Println("Got ping tx request")
	return ping.HandlePingTransactionRequest(req)
}
