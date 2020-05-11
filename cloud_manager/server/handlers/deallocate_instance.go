package handlers

import (
	"log"

	cloud "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/api/proto/common"
)

// See [Go SDK] https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.ReleaseHosts
//     [API]    https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReleaseHosts.html
func HandleDeallocateVMRequest(req *cloud.TDeallocateInstanceRequest) cloud.TDeallocateInstanceResponse {
	log.Printf("got Deallocate Instance req: %s", req.InstanceUuid)
	return cloud.TDeallocateInstanceResponse{
		Result: &common.TResult{ResultCode: common.TResult_FAILED, ErrorCode: common.TResult_NOT_IMPLEMENTED},
	}
}
