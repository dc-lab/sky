package async

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

func HandleStartTaskRequest(req *cm.TStartTaskRequest) (resp *cm.TAsyncCloudResponse, err error) {
	log.Printf("got Start Task req for %s from %s", req.GetFactoryId(), req.GetUserId())
	return nil, status.Error(codes.Unimplemented, "Start Task is not supported")
}
