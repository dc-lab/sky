package async

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

// See [Go SDK] https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.ReleaseHosts
//     [API]    https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReleaseHosts.html
func HandleDeallocateInstanceRequest(req *cm.TDeallocateInstanceRequest) (resp *cm.TAsyncCloudResponse, err error) {
	log.Printf("got Deallocate Instance req for instance %s from %s", req.GetInstanceId(), req.GetUserId())
	return nil, status.Error(codes.Unimplemented, "Deallocate Instance is not supported")
}
