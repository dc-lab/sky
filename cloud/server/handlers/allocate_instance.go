package handlers

import (
	"log"

	"github.com/dc-lab/sky/api/proto/cloud"
	"github.com/dc-lab/sky/api/proto/common"
)

// See [Go SDK] https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.AllocateHosts
//     [API]    https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllocateHosts.html
func HandleAllocateVMRequest(req *cloud.TAllocateInstanceRequest) cloud.TAllocateInstanceResponse {
	log.Printf("got Allocate Instance req: %s", req.HardwareData)
	return cloud.TAllocateInstanceResponse{
		Result: &common.TResult{ResultCode: common.TResult_FAILED, ErrorCode: common.TResult_NOT_IMPLEMENTED},
	}
}
