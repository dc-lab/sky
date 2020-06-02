package async

import (
	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleDeployImageRequest(req *cm.TDeployImageRequest) (resp *cm.TAsyncCloudResponse, err error) {
	log.Printf("got Deploy Image req for instance %s and image %s from %s", req.GetInstanceId(), req.GetDockerImage(), req.GetUserId())
	return nil, status.Error(codes.Unimplemented, "Deploy Image is not supported")
}
