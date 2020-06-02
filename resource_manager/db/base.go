package db

import (
	"context"
	"fmt"
	"github.com/dc-lab/sky/resource_manager/app"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
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
	username := app.Config.DBUser
	password := os.Getenv(app.Config.DBPasswordEnv)
	host := app.Config.DBHost
	dbName := app.Config.DBName
	ssl := app.Config.DBSsl
	dbUri := fmt.Sprintf("postgres://%s:%s@%s/%s?ssl=%v", username, password, host, dbName, ssl)

	var err error
	pool, err = pgxpool.Connect(context.Background(), dbUri)
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
