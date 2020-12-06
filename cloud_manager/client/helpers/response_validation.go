package helpers

import (
	"log"

	"github.com/dc-lab/sky/api/proto/common"
)

type WithResult interface {
	GetResult() *common.TResult
}

func EnsureOkStatusCode(obj WithResult) {
	code := obj.GetResult().GetResultCode()
	if code != common.TResult_SUCCESS {
		log.Fatalf("unexpected result code: %v", code)
	}
}
