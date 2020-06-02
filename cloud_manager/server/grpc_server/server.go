package grpc_server

import (
	"context"
	"github.com/dc-lab/sky/cloud_manager/server/db"
	"github.com/dc-lab/sky/cloud_manager/server/grpc_client"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/cloud_manager/server/grpc_server/async"
	"github.com/dc-lab/sky/cloud_manager/server/grpc_server/ping"
	"github.com/dc-lab/sky/cloud_manager/server/grpc_server/transaction"
)

type scope struct {
	tx *transaction.Tx
}

type Server struct{
	txManager *transaction.TxManager
	rmClient *grpc_client.ResourceManagerClient
}

func New(dbClient *db.Client, rmClient *grpc_client.ResourceManagerClient) *Server {
	return &Server{
		txManager: transaction.NewTxManager(dbClient),
		rmClient: rmClient,
	}
}

func (s Server) DoAsync(_ context.Context, req *cm.TAsyncCloudRequest) (*cm.TAsyncCloudResponse, error) {
	log.Println("Got async cloud request")

	var resp *cm.TAsyncResponseContent
	switch typedReq := req.Body.(type) {
	case *cm.TAsyncCloudRequest_AllocateInstanceRequest:
		typedResp, err := async.HandleAllocateInstanceRequest(typedReq.AllocateInstanceRequest)
		if err != nil {
			return nil, err
		}
		resp = &cm.TAsyncResponseContent{
			Body: &cm.TAsyncResponseContent_AllocateInstanceResponseContent{
				AllocateInstanceResponseContent: typedResp,
			},
		}
	case *cm.TAsyncCloudRequest_DeallocateInstanceRequest:
		typedResp, err := async.HandleDeallocateInstanceRequest(typedReq.DeallocateInstanceRequest)
		if err != nil {
			return nil, err
		}
		resp = &cm.TAsyncResponseContent{
			Body: &cm.TAsyncResponseContent_DeallocateInstanceResponseContent{
				DeallocateInstanceResponseContent: typedResp,
			},
		}
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

	log.Println("Send cloud response")
	tx := s.txManager.NewExternalOp("Handle request")
	return &cm.TAsyncCloudResponse{
		Transaction:          tx.ToPB(),
		AsyncResponseContent: resp,
	}, nil
}

func (s Server) PingTransaction(_ context.Context, req *cm.TPingTransactionRequest) (*cm.TPingTransactionResponse, error) {
	log.Println("Got ping tx request")
	res, err := ping.HandlePingTransactionRequest(req, s.txManager)
	if err != nil {
		return nil, err
	}
	return res, nil
}
