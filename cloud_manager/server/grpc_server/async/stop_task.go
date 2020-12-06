package async

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

func HandleStopTaskRequest(req *cm.TStopTaskRequest) (resp *cm.TAsyncCloudResponse, err error) {
	log.Printf("got Stop Task req for %s from %s", req.GetInstanceId(), req.GetUserId())
	awsInp := "{\n    \"InterruptTask\": [\n        {\n            \"TaskId\": \"%s\"}]\n}"
	awsInp = fmt.Sprintf(awsInp, req.GetInstanceId())
	return nil, status.Error(codes.Unimplemented, "Stop Task is not supported")
}