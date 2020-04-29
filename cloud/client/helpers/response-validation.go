package helpers

import (
	"log"

	pb "github.com/dc-lab/sky/cloud/proto"
)

type WithResult interface {
	GetResult() *pb.TResult
}

func EnsureOkStatusCode(obj WithResult) {
	code := obj.GetResult().GetResultCode()
	if code != pb.TResult_SUCCESS {
		log.Fatalf("unexpected result code: %v", code)
	}
}
