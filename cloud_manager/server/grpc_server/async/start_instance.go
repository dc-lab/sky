package async

import (
	"fmt"
	"log"
	"aws"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

func HandleStartInstanceRequest(req *cm.TStartInstanceRequest) (resp *cm.TAsyncCloudResponse, err error) {
	log.Printf("got Start Instance req for %s from %s", req.GetInstanceId(), req.GetUserId())
	awsInp := "{\n    \"StartingInstances\": [\n        {\n            \"InstanceId\": \"%s\",\n            \"CurrentState\": {\n                \"Code\": 0,\n                \"Name\": \"pending\"\n            },\n            \"PreviousState\": {\n                \"Code\": 80,\n                \"Name\": \"stopped\"\n            }\n        }\n    ]\n}"
	awsInp = fmt.Sprintf(awsInp, req.GetInstanceId())
	aws.startInstances(awsInp)
	return nil, status.Error(codes.Unimplemented, "Start Instance is not supported")
}
