package grpc_server

import (
	"log"

	cloud "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/api/proto/common"
)

func HandleDisconnectInstanceRequest(req *cloud.TDisconnectInstanceRequest) cloud.TDisconnectInstanceResponse {
	log.Printf("got Disconnect Instance req: %s", req.InstanceUuid)
	return cloud.TDisconnectInstanceResponse{
		Result: &common.TResult{ResultCode: common.TResult_FAILED, ErrorCode: common.TResult_NOT_IMPLEMENTED},
	}
}
