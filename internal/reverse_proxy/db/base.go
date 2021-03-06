package db

import (
	"context"
	"fmt"
	"github.com/dc-lab/sky/internal/reverse_proxy/app"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"os"
)

var pool *pgxpool.Pool

func InitDB() {
	username := app.Config.DBUser
	password := os.Getenv(app.Config.DBPasswordEnv)
	host := app.Config.DBHost
	dbName := app.Config.DBName
	ssl := "disable"
	if app.Config.DBSsl {
		ssl = "require"
	}
	dbUri := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%v", username, password, host, dbName, ssl)

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

	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users (id varchar(40) PRIMARY KEY, login varchar(256) NOT NULL UNIQUE, password varchar(256) NOT NULL, token varchar(256) NOT NULL);")
	if err != nil {
		log.Fatal(err)
	}
}
