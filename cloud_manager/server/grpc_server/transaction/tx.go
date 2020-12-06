package transaction

import (
	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
	"time"
)

type Tx struct {
	id string
	externalId *string
	expireAt time.Time
	status int
}

const (
	statusActive int = 1
	statusCompleted int = 2
	statusExpired int = 3
	statusCancelled int = 4
)

var status2pb = map[int]cm.ETransactionStatus{
	statusActive: cm.ETransactionStatus_ACTIVE,
	statusCompleted: cm.ETransactionStatus_COMPLETED,
	statusExpired: cm.ETransactionStatus_EXPIRED,
	statusCancelled: cm.ETransactionStatus_CANCELLED,
}

func (tx *Tx) isActive() bool {
	return tx.status == statusActive
}

func (tx *Tx) isExpired() bool {
	return tx.status == statusExpired
}

func (tx *Tx) PbTx() *cm.TTransaction {
	return &cm.TTransaction{
		TransactionId: tx.id,
	}
}

func (tx *Tx) PbStatus() cm.ETransactionStatus {
	if res, exist := status2pb[tx.status]; exist {
		return res
	}
	return cm.ETransactionStatus_UNKNOWN
}
