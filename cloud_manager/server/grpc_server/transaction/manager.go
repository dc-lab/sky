package transaction

import (
	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/cloud_manager/server/db"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type TxManager struct {
	dao *db.PsqlTxDao
}

const (
	DefaultExternalTtl = 10 * time.Minute
	DefaultInternalTtl = 3 * time.Second
)

func NewTxManager(dbClient *db.Client) *TxManager {
	dao := db.NewPsqlTxDao(dbClient)
	return &TxManager{dao}
}

func (m *TxManager) NewExternalTx() (*Tx, error) {
	return m.newTxWithTtl(DefaultExternalTtl)
}

func (m *TxManager) NewInternalTx() (*Tx, error) {
	return m.newTxWithTtl(DefaultInternalTtl)
}

func (m *TxManager) AddInternalOp(tx Tx) (*Tx, error) {
	return m.AddOp(tx, DefaultInternalTtl)
}

func (m *TxManager) AddExternalOp(tx Tx) (*Tx, error) {
	return m.AddOp(tx, DefaultInternalTtl)
}

func (m *TxManager) GetStatus(txId string) (cm.ETransactionStatus, error) {
	tx := (*m.dao).Get(txId)
	if tx == nil {
		return cm.ETransactionStatus_UNKNOWN, status.Error(codes.NotFound, "Not found tx " + txId)
	}
	return tx.PbStatus(), nil
}

func (m *TxManager) newTxWithTtl(ttl time.Duration) (*Tx, error) {
	id := uuid.New().String()
	expireAt := time.Now().Add(ttl)
	tx := &Tx{
		id:  id,
		expireAt: expireAt,
		status: statusActive,
	}
	if err := (*m.dao).Insert(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

func (m *TxManager) extendTtl(tx *Tx) error {
	if !tx.isActive() {
		return xerrors.New("Tx is inactive")
	}
	tx.expireAt = tx.expireAt.Add(DefaultInternalTtl)
	if err := (*m.dao).Update(tx); err != nil {
		return err
	}
	return nil
}