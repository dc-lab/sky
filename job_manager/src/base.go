package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

var pool *pgxpool.Pool

func getEnv(key, fallback string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	return fallback
}

func _() {
	username := getEnv("DB_USER", "oleg")
	password := os.Getenv("DB_PASSWORD")
	host := getEnv("DB_HOST", "rc1b-6marivlovkr6pccx.mdb.yandexcloud.net:6432")
	dbName := getEnv("DB_NAME", "sky_postgre")
	ssl := getEnv("DB_SSL", "true")
	dbUri := fmt.Sprintf("postgres://%s:%s@%s/%s?ssl=%s", username, password, host, dbName, ssl)

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

	//_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users (id varchar(40) PRIMARY KEY, login varchar(256) NOT NULL UNIQUE, password varchar(256) NOT NULL, token varchar(256) NOT NULL);")
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func _() *pgxpool.Pool {
	return pool
}
