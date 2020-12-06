package async

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

func HandleStopInstanceRequest(req *cm.TStopInstanceRequest) (resp *cm.TAsyncCloudResponse, err error) {
	log.Printf("got Stop Instance req for %s from %s", req.GetInstanceId(), req.GetUserId())
	awsInp := "{\n    \"StoppingInstances\": [\n        {\n            \"InstanceId\": \"%s\",\n            \"CurrentState\": {\n                \"Code\": 0,\n                \"Name\": \"running\"\n            },\n            \"PreviousState\": {\n                \"Code\": 80,\n                \"Name\": \"starting\"\n            }\n        }\n    ]\n}"
	awsInp = fmt.Sprintf(awsInp, req.GetInstanceId())
	aws.stopInstances(awsInp)
	return nil, status.Error(codes.Unimplemented, "Stop Instance is not supported")
}
