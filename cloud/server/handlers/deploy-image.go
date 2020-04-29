package handlers

import (
	"log"

	pb "github.com/dc-lab/sky/cloud/proto"
)

func HandleDeployImageRequest(req *pb.TDeployImageRequest) pb.TDeployImageResponse {
	log.Printf("got Deploy Image req: %s, %s", req.InstanceUuid, req.DockerImage)
	return pb.TDeployImageResponse{
		Result: &pb.TResult{ResultCode: pb.TResult_FAILED, ErrorCode: pb.TResult_NOT_IMPLEMENTED},
	}
}
