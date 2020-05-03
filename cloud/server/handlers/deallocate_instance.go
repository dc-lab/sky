package handlers

import (
	"log"

	"github.com/dc-lab/sky/api/proto/cloud"
	"github.com/dc-lab/sky/api/proto/common"
)

func HandleDeallocateVMRequest(req *cloud.TDeallocateInstanceRequest) cloud.TDeallocateInstanceResponse {
	log.Printf("got Deallocate Instance req: %s", req.InstanceUuid)
	return cloud.TDeallocateInstanceResponse{
		Result: &common.TResult{ResultCode: common.TResult_FAILED, ErrorCode: common.TResult_NOT_IMPLEMENTED},
	}
}
