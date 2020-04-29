package handlers

import (
	"log"

	pb "github.com/dc-lab/sky/cloud/proto"
)

func HandleStartVMRequest(req *pb.TStartInstanceRequest) pb.TStartInstanceResponse {
	log.Printf("got Start Instance req: %s", req.InstanceUuid)
	return pb.TStartInstanceResponse{
		Result: &pb.TResult{ResultCode: pb.TResult_FAILED, ErrorCode: pb.TResult_NOT_IMPLEMENTED},
	}
}
