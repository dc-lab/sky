package handlers

import (
	"log"

	cloud "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/api/proto/common"
)

func HandleStopVMRequest(req *cloud.TStopInstanceRequest) cloud.TStopInstanceResponse {
	log.Printf("got Stop Instance req: %s", req.InstanceUuid)
	return cloud.TStopInstanceResponse{
		Result: &common.TResult{ResultCode: common.TResult_FAILED, ErrorCode: common.TResult_NOT_IMPLEMENTED},
	}
}
