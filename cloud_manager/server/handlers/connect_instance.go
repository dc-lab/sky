package handlers

import (
	"log"

	cloud "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/api/proto/common"
)

// See [Go SDK] https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.DescribeHosts
//     [API]    https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeHosts.html
func HandleConnectInstanceRequest(req *cloud.TConnectInstanceRequest) cloud.TConnectInstanceResponse {
	log.Printf("got Connect Instance req: %s", req.HardwareData)
	return cloud.TConnectInstanceResponse{
		Result: &common.TResult{ResultCode: common.TResult_FAILED, ErrorCode: common.TResult_NOT_IMPLEMENTED},
	}
}
