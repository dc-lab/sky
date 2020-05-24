package async

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

func HandleStartInstanceRequest(req *cm.TStartInstanceRequest) (resp *cm.TAsyncCloudResponse, err error) {
	log.Printf("got Start Instance req for %s from %s", req.GetInstanceId(), req.GetUserId())
	return nil, status.Error(codes.Unimplemented, "Start Instance is not supported")
}
