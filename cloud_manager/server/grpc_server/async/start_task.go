package async

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

func HandleStartTaskRequest(req *cm.TStartTaskRequest) (resp *cm.TAsyncCloudResponse, err error) {
	log.Printf("got Start Task req for %s from %s", req.GetFactoryId(), req.GetUserId())
	awsInp := "{\n    \"StartTask\": [\n        {\n            \"TaskId\": \"%s\",\n            \"CurrentState\": {\n                \"Code\": 0,\n                \"Name\": \"pending\"\n            }]\n}"
	awsInp = fmt.Sprintf(awsInp, req.GetFactoryId())
	return nil, status.Error(codes.Unimplemented, "Start Task is not supported")
}
