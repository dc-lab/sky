package handlers

import (
	"log"

	cloud "github.com/dc-lab/sky/cloud/proto"
)

func HandleStopVMRequest(req *cloud.TStopInstanceRequest) cloud.TStopInstanceResponse {
	log.Printf("got Stop Instance req: %s", req.InstanceUuid)
	return cloud.TStopInstanceResponse{
		Result: &cloud.TResult{ResultCode: cloud.TResult_FAILED, ErrorCode: cloud.TResult_NOT_IMPLEMENTED},
	}
}
