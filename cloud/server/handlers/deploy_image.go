package handlers

import (
	"log"

	"github.com/dc-lab/sky/api/proto/cloud"
	"github.com/dc-lab/sky/api/proto/common"
)

func HandleDeployImageRequest(req *cloud.TDeployImageRequest) cloud.TDeployImageResponse {
	log.Printf("got Deploy Image req: %s, %s", req.InstanceUuid, req.DockerImage)
	return cloud.TDeployImageResponse{
		Result: &common.TResult{ResultCode: common.TResult_FAILED, ErrorCode: common.TResult_NOT_IMPLEMENTED},
	}
}
