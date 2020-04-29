package handlers

import (
	"log"

	pb "github.com/dc-lab/sky/cloud/proto"
)

func HandleDisconnectInstanceRequest(req *pb.TDisconnectInstanceRequest) pb.TDisconnectInstanceResponse {
	log.Printf("got Disconnect Instance req: %s", req.InstanceUuid)
	return pb.TDisconnectInstanceResponse{
		Result: &pb.TResult{ResultCode: pb.TResult_FAILED, ErrorCode: pb.TResult_NOT_IMPLEMENTED},
	}
}
