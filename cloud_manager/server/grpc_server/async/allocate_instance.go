package async

import (
	"github.com/aws/aws-sdk-go/aws"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

// See [Go SDK] https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.AllocateHosts
//     [API]    https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllocateHosts.html
func HandleAllocateInstanceRequest(req *cm.TAllocateInstanceRequest) (resp *cm.TAllocateInstanceResponseContent, err error) {
	log.Printf("got Allocate Instance req for factory %s from %s", req.GetFactoryId(), req.GetUserId())
	aws.allocate(InstanceIds: req.FactoryId, UserId: req.UserId)
	return nil, status.Error(codes.Unimplemented, "Allocate Instance is not supported")
}
