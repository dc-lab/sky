package handlers

import (
	"log"

	pb "github.com/dc-lab/sky/cloud/proto"
)

func HandleDeallocateVMRequest(req *pb.TDeallocateInstanceRequest) pb.TDeallocateInstanceResponse {
	log.Printf("got Deallocate Instance req: %s", req.InstanceUuid)
	return pb.TDeallocateInstanceResponse{
		Result: &pb.TResult{ResultCode: pb.TResult_FAILED, ErrorCode: pb.TResult_NOT_IMPLEMENTED},
	}
}
