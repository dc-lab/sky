package db

import (
	"context"
	"github.com/dc-lab/sky/internal/resource_manager/app"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

var pool *pgxpool.Pool

func initTables(conn *pgxpool.Conn) {
	_, err := conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS resources (id varchar(40) PRIMARY KEY, owner_id varchar(40) NOT NULL, name varchar(256) NOT NULL, type varchar(20) NOT NULL, token varchar(40) NOT NULL UNIQUE);")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS ur_permissions (user_id varchar(40), resource_id varchar(40) REFERENCES resources ON DELETE CASCADE);")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS gr_permissions (group_id varchar(40), resource_id varchar(40) REFERENCES resources ON DELETE CASCADE);")
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB() {
	var err error
	pool, err = pgxpool.Connect(context.Background(), app.Config.PostgresAddress)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Can't perform basic db check on start: %v", err)
	}
	defer conn.Release()

	initTables(conn)
}

func GetPool() *pgxpool.Pool {
	return pool
}
