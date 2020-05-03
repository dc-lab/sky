package handlers

import (
	"log"

	"github.com/dc-lab/sky/api/proto/cloud"
	"github.com/dc-lab/sky/api/proto/common"
)

func HandleConnectInstanceRequest(req *cloud.TConnectInstanceRequest) cloud.TConnectInstanceResponse {
	log.Printf("got Connect Instance req: %s", req.HardwareData)
	return cloud.TConnectInstanceResponse{
		Result: &common.TResult{ResultCode: common.TResult_FAILED, ErrorCode: common.TResult_NOT_IMPLEMENTED},
	}
}
