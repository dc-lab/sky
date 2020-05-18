package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc"

	cm "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/cloud_manager/server/app"
	"github.com/dc-lab/sky/cloud_manager/server/cmd"
	"github.com/dc-lab/sky/cloud_manager/server/common"
	"github.com/dc-lab/sky/cloud_manager/server/db"
	"github.com/dc-lab/sky/cloud_manager/server/grpc_server"
)

func gRPCStarter(wg *sync.WaitGroup, address string) {
	defer wg.Done()

	// create listener
	log.Printf("Start listening tcp on %d", cmd.GrpcPort)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", address))
	common.DieWithError("Failed to listen:", err)

	// create gRPC server
	s := grpc.NewServer()
	cm.RegisterTCloudManagerServer(s, grpc_server.Server{})

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

func main() {
	app.ParseConfig()

	logFile := initLogs(app.Config.LogsFile)
	defer logFile.Close()

	db.InitDB()

	var wg sync.WaitGroup

	wg.Add(1)
	go gRPCStarter(&wg, app.Config.GRPCAddress)

	wg.Wait()
}
