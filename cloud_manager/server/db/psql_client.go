package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Conn *sql.DB
}

func OpenConnection(connStr string, applyMigrations bool) (*Client, error) {
	conn, err := sql.Open("postgres", connStr)
	if err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()

		pingOperation := func () error { return conn.PingContext(ctx) }
		err = backoff.Retry(pingOperation, backoff.NewExponentialBackOff())
	}

	client := &Client{
		Conn: conn,
	}

	if err != nil {
		return client, err
	}

	if applyMigrations {
		err = client.migrate()
	}

	if err != nil {
		log.WithError(err).Fatalln("Failed to run migrations")
	}

	return client, err
}

func (c *Client) migrate() error {
	driver, err := postgres.WithInstance(c.Conn, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		err = nil
	}

	return err
}
