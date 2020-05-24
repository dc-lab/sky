package async

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

func HandleStopInstanceRequest(req *cm.TStopInstanceRequest) (resp *cm.TAsyncCloudResponse, err error) {
	log.Printf("got Stop Instance req for %s from %s", req.GetInstanceId(), req.GetUserId())
	return nil, status.Error(codes.Unimplemented, "Stop Instance is not supported")
}
