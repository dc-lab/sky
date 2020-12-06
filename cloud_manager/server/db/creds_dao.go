package db

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
)

type PsqlCredsDao struct {
	dbClient     *Client
	userTimeout time.Duration
	readTimeout time.Duration
}

func NewPsqlCredsDao(dbClient *Client) *PsqlCredsDao {
	return &PsqlCredsDao{
		dbClient:     dbClient,
		userTimeout: defaultOfflineTimeout,
		readTimeout: defaultRuntimeReadTimeout,
	}
}

func (d *PsqlCredsDao) Create(creds *Credentials) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.userTimeout)
	defer cancel()

	err := d.dbClient.Conn.QueryRowContext(
		ctx,
		`INSERT INTO credentials (
			owner_id,
			display_name,
			provider,
			aws_access_key_id,
			aws_access_key
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id`,
		creds.OwnerId,
		creds.Provider,
		creds.AwsAccessKeyId,
		creds.AwsAccessKey,
	).Scan(&creds.Id)
	return err
}

func (d *PsqlCredsDao) GetAll(ownerId string) ([]Credentials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), d.userTimeout)
	defer cancel()

	rows, err := d.dbClient.Conn.QueryContext(
		ctx,
		`SELECT
			id,
			owner_id,
			display_name,
			provider,
			aws_access_key_id,
			aws_access_key
		FROM credentials
		WHERE owner_id=$1`,
		ownerId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	allCreds := make([]Credentials, 0)

	for rows.Next() {
		var creds Credentials
		err := rows.Scan(
			&creds.Id,
			&creds.OwnerId,
			&creds.DisplayName,
			&creds.Provider,
			&creds.AwsAccessKeyId,
			&creds.AwsAccessKey)
		if err != nil {
			log.Tracef("Failed to get credentials row. Skipping... %v", err)
			continue
		}
		allCreds = append(allCreds, creds)
	}

	return allCreds, nil
}

func (d *PsqlCredsDao) Get(ownerId string, credsId string) (*Credentials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), d.readTimeout)
	defer cancel()

	var creds Credentials
	err := d.dbClient.Conn.QueryRowContext(
		ctx,
		`SELECT
			id,
			owner_id,
			display_name,
			provider,
			aws_access_key_id,
			aws_access_key
		FROM credentials
		WHERE id=$1 AND owner_id=$2`,
		credsId,
		ownerId,
	).Scan(&creds.Id, &creds.OwnerId, &creds.DisplayName, &creds.Provider, &creds.AwsAccessKeyId, &creds.AwsAccessKey)
	return &creds, err
}

func (d *PsqlCredsDao) Delete(ownerId string, credsId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.userTimeout)
	defer cancel()

	_, err := d.dbClient.Conn.ExecContext(
		ctx,
		`DELETE FROM credentials
		WHERE id=$1 AND owner_id=$2`,
		credsId,
		ownerId,
	)
	return err
}
