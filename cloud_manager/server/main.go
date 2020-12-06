package main

import (
	"fmt"
	"github.com/dc-lab/sky/cloud_manager/server/grpc_client"
	"github.com/dc-lab/sky/cloud_manager/server/http_handles"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"google.golang.org/grpc"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/cloud_manager/server/app"
	"github.com/dc-lab/sky/cloud_manager/server/common"
	"github.com/dc-lab/sky/cloud_manager/server/db"
	"github.com/dc-lab/sky/cloud_manager/server/grpc_server"
)

func httpStarter(wg *sync.WaitGroup, address string, dp *http_handles.DataProvider) {
	defer wg.Done()

	router := http_handles.NewRouter(dp)

	log.Printf("Start listening http on %s", address)
	err := http.ListenAndServe(address, router)
	common.DieWithError("Failed to listen HTTP:", err)
}

func gRPCStarter(wg *sync.WaitGroup, address string, dbClient *db.Client, rmClient *grpc_client.ResourceManagerClient) {
	defer wg.Done()

	// create listener
	log.Printf("Start listening tcp on %s", address)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", address))
	common.DieWithError("Failed to listen gRPC:", err)

	// create gRPC server
	s := grpc.NewServer()
	srv := grpc_server.New(dbClient, rmClient)
	cm.RegisterTCloudManagerServer(s, srv)

	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func initLogs(logFilePath string) *os.File {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	common.DieWithError("Cannot open logs file: " + logFilePath, err)
	log.SetOutput(file)
	return file
}

func makeConnStr() string {
	username := app.Config.DBUser
	password := os.Getenv(app.Config.DBPasswordEnv)
	host := app.Config.DBHost
	dbName := app.Config.DBName
	ssl := app.Config.DBSsl
	return fmt.Sprintf("postgres://%s:%s@%s/%s?ssl=%v", username, password, host, dbName, ssl)
}

func main() {
	app.ParseConfig()

	logFile := initLogs(app.Config.LogsFile)
	defer logFile.Close()

	dbClient, err := db.OpenConnection(makeConnStr(), app.Config.ApplyMigrationsOnStart)
	common.DieWithError("Cannot open connection with db", err)
	defer dbClient.Conn.Close()

	rmClient := grpc_client.NewResourceManagerClient(app.Config.ResourceManagerAddress)

	var wg sync.WaitGroup

	wg.Add(1)
	go httpStarter(&wg, app.Config.HTTPAddress)

	wg.Add(1)
	go gRPCStarter(&wg, app.Config.GRPCAddress, dbClient, rmClient)

	wg.Wait()
}
