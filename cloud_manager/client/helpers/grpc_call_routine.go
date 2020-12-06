package helpers

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	cloud "github.com/dc-lab/sky/api/proto/cloud_manager"
)

func MakeCloudRequest(grpcPort uint16, req *cloud.TCloudRequest) *cloud.TCloudResponse {
	log.Println("connect called")

	conn, err := grpc.Dial(fmt.Sprintf(":%d", grpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect to server %v", err)
	}

	// create stream
	client := cloud.NewTCloudManagerClient(conn)
	stream, err := client.DoAction(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	// send message
	if err := stream.Send(req); err != nil {
		log.Fatalf("send request error %v", err)
	}

	// receive message
	resp, err := stream.Recv()
	if err != nil {
		log.Fatalf("recv response error %v", err)
	}

	return resp
}
