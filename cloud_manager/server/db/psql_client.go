package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/dc-lab/sky/cloud_manager/server/app"
	"github.com/dc-lab/sky/cloud_manager/server/common"
)

var pool *pgxpool.Pool

// TODO: migrations
func InitDB() {
	username := app.Config.DBUser
	password := os.Getenv(app.Config.DBPasswordEnv)
	host := app.Config.DBHost
	dbName := app.Config.DBName
	ssl := app.Config.DBSsl
	dbUri := fmt.Sprintf("postgres://%s:%s@%s/%s?ssl=%v", username, password, host, dbName, ssl)

	pool, err := pgxpool.Connect(context.Background(), dbUri)
	common.DieWithError("Cannot connect to psql", err)

	conn, err := pool.Acquire(context.Background())
	common.DieWithError("Cannot perform basic db check on start", err)
	defer conn.Release()

	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS credentials (id varchar(40) PRIMARY KEY, owner_id varchar(40) NOT NULL, aws_access_key_id varchar(128) NOT NULL, aws_access_key varchar(128) NOT NULL);")
	common.DieWithError("Unexpected error during db startup", err)
}

func GetPool() *pgxpool.Pool {
	return pool
}
