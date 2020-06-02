package grpc_client

import (
	"context"
	rm "github.com/dc-lab/sky/api/proto/resource_manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type ResourceManagerClient struct {
	resourceManagerAddr string
	requestTimeout time.Duration
}

const DefaultRequestTimeout = 1 * time.Second

func NewResourceManagerClient(resourceManagerAddr string) *ResourceManagerClient {
	return &ResourceManagerClient{
		resourceManagerAddr: resourceManagerAddr,
		requestTimeout: DefaultRequestTimeout,
	}
}

func (c *ResourceManagerClient) CreateResource(resource *rm.TResource) error {
	req := &rm.TResourceRequest{
		Body: &rm.TResourceRequest_CreateResourceRequest{
			CreateResourceRequest: &rm.TCreateResourceRequest{
				Resource: resource,
			},
		},
	}
	_, err := c.update(req)
	return err
}

func (c *ResourceManagerClient) DeleteResource(resourceId string, userId string) error {
	req := &rm.TResourceRequest{
		Body: &rm.TResourceRequest_DeleteResourceRequest{
			DeleteResourceRequest: &rm.TDeleteResourceRequest{
				ResourceId: resourceId,
				UserId: userId,
			},
		},
	}
	_, err := c.update(req)
	return err
}

func (c *ResourceManagerClient) update(req *rm.TResourceRequest) (*rm.TResourceResponse, error) {
	conn, err := grpc.Dial(c.resourceManagerAddr)
	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot connect with RM on addr " + c.resourceManagerAddr)
	}
	defer conn.Close()

	grpcClient := rm.NewResourceManagerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()

	return grpcClient.Update(ctx, req)
}
