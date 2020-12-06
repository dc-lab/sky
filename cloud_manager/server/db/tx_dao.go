package db

import (
	"context"
	"time"
)

type PsqlTxDao struct {
	dbClient     *Client
	writeTimeout time.Duration
	readTimeout  time.Duration
}

func NewPsqlTxDao(dbClient *Client) *PsqlTxDao {
	return &PsqlTxDao{
		dbClient:     dbClient,
		writeTimeout: defaultRuntimeWriteTimeout,
		readTimeout:  defaultRuntimeReadTimeout,
	}
}

func (d *PsqlTxDao) Create(tx *Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.writeTimeout)
	defer cancel()

	err := d.dbClient.Conn.QueryRowContext(
		ctx,
		`INSERT INTO transactions (
			external_id,
			expire_at,
			status,
		)
		VALUES ($1, $2, $3)
		RETURNING id`,
		tx.ExternalId,
		tx.ExpireAt,
		tx.Status,
	).Scan(&tx.Id)
	return err
}

func (d *PsqlTxDao) Update(tx *Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.writeTimeout)
	defer cancel()

	_, err := d.dbClient.Conn.ExecContext(
		ctx,
		`UPDATE transactions
		SET
			external_id=$2,
			expire_at=$3,
			status=$4
		WHERE id=$1`,
		tx.Id,
		tx.ExternalId,
		tx.ExpireAt,
		tx.Status,
	)
	return err
}

func (d *PsqlTxDao) Get(txId string) (*Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), d.readTimeout)
	defer cancel()

	var tx Transaction
	err := d.dbClient.Conn.QueryRowContext(
		ctx,
		`SELECT
			id,
			external_id,
			expire_at,
			status
		FROM transactions
		WHERE id=$1`,
		txId,
	).Scan(&tx.Id, &tx.ExternalId, &tx.ExpireAt, &tx.Status)
	return &tx, err
}

func (d *PsqlTxDao) Delete(txId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.writeTimeout)
	defer cancel()

	_, err := d.dbClient.Conn.ExecContext(
		ctx,
		`DELETE FROM transactions
		WHERE id=$1`,
		txId,
	)
	return err
}
