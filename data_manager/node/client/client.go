package client

import (
	"context"
	"time"

	"google.golang.org/grpc"

	pb "github.com/dc-lab/sky/api/proto/data_manager"
	"github.com/dc-lab/sky/data_manager/node/config"
)

type Client struct {
	config *config.Config
	nodeId string
	client pb.MasterClient
}

func MakeClient(config *config.Config) (*Client, error) {
	conn, err := grpc.Dial(config.MasterAddress, grpc.WithInsecure(), grpc.WithBackoffMaxDelay(time.Second*10))
	if err != nil {
		return nil, err
	}
	client := pb.NewMasterClient(conn)
	return &Client{config, config.AccessAddress, client}, nil
}

func (c *Client) Loop(state *pb.NodeStatus) (*pb.NodeTarget, error) {
	state.NodeId = c.nodeId
	target, err := c.client.Loop(context.Background(), state)
	return target, err
}

func (c *Client) ValidateUpload(user_id string, file_id string, token string) (bool, error) {
	res, err := c.client.ValidateUpload(context.Background(), &pb.ValidateUploadRequest{
		NodeId:      c.config.AccessAddress,
		UserId:      user_id,
		FileId:      file_id,
		UploadToken: token,
	})

	return res.GetAllow(), err
}

func (c *Client) SubmitFileHash(user_id string, file_id string, hash string) (bool, error) {
	res, err := c.client.SubmitFileHash(context.Background(), &pb.SubmitFileHashRequest{
		NodeId: c.config.AccessAddress,
		FileId: file_id,
		UserId: user_id,
		Hash:   hash,
	})

	return res.GetAllow(), err
}

func (c *Client) GetFileHash(file_id string, user_id string) (bool, string, error) {
	res, err := c.client.GetFileHash(context.Background(), &pb.GetFileHashRequest{
		NodeId: c.config.AccessAddress,
		UserId: user_id,
		FileId: file_id,
	})

	return res.GetAllow(), res.GetHash(), err
}
