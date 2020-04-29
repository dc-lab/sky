package handlers

import (
	"log"

	pb "github.com/dc-lab/sky/cloud/proto"
)

func HandleConnectInstanceRequest(req *pb.TConnectInstanceRequest) pb.TConnectInstanceResponse {
	log.Printf("got Connect Instance req: %s", req.HardwareData)
	return pb.TConnectInstanceResponse{
		Result: &pb.TResult{ResultCode: pb.TResult_FAILED, ErrorCode: pb.TResult_NOT_IMPLEMENTED},
	}
}
