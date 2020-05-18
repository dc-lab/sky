package grpc_server

import (
	"log"

	cloud "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/api/proto/common"
)

func HandleStartVMRequest(req *cloud.TStartInstanceRequest) cloud.TStartInstanceResponse {
	log.Printf("got Start Instance req: %s", req.InstanceUuid)
	return cloud.TStartInstanceResponse{
		Result: &common.TResult{ResultCode: common.TResult_FAILED, ErrorCode: common.TResult_NOT_IMPLEMENTED},
	}
}
