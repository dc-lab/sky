package ping

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
)

func HandlePingTransactionRequest(req *cm.TPingTransactionRequest) (resp *cm.TPingTransactionResponse, err error) {
	log.Printf("got ping tx req for tx %s", req.GetTransactionId())
	return nil, status.Error(codes.Unimplemented, "Ping Transaction is not supported")
}