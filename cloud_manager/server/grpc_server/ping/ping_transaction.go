package ping

import (
	"log"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/cloud_manager/server/grpc_server/transaction"
)

func HandlePingTransactionRequest(req *cm.TPingTransactionRequest, txManager *transaction.TxManager) (resp *cm.TPingTransactionResponse, err error) {
	log.Printf("got ping tx req for tx %s", req.GetTransactionId())
	txStatus, err := txManager.GetStatus(req.GetTransactionId())
	if err != nil {
		return nil, err
	}
	return &cm.TPingTransactionResponse{
		Status: txStatus,
	}, nil
}
