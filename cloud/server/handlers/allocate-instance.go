package handlers

import (
	"log"

	pb "github.com/dc-lab/sky/cloud/proto"
)

func HandleAllocateVMRequest(req *pb.TAllocateInstanceRequest) pb.TAllocateInstanceResponse {
	log.Printf("got Allocate Instance req: %s", req.HardwareData)
	return pb.TAllocateInstanceResponse{
		Result: &pb.TResult{ResultCode: pb.TResult_FAILED, ErrorCode: pb.TResult_NOT_IMPLEMENTED},
	}
}
