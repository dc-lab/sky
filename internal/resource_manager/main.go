package main

import (
	pb "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/resource_manager/app"
	"github.com/dc-lab/sky/internal/resource_manager/db"
	"github.com/dc-lab/sky/internal/resource_manager/grpc_server"
	"github.com/dc-lab/sky/internal/resource_manager/http_handles"
	"github.com/dc-lab/sky/pkg/logging"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"sync"
)

func httpStarter(wg *sync.WaitGroup, addr string) {
	defer wg.Done()
	router := mux.NewRouter()

	router.HandleFunc("/resources", http_handles.Resources).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/resources/{id}", http_handles.Resource).Methods(http.MethodGet, http.MethodDelete, http.MethodPost)

	log.Info("Starting http server")

	log.Fatal(http.ListenAndServe(addr, router))
}

func gRPCStarter(wg *sync.WaitGroup, addr string, dmAddress string) {
	defer wg.Done()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server, err := grpc_server.NewServer(dmAddress)
	if err != nil {
		log.Fatalf("Failed to setup connection with data manager: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterResourceManagerServer(s, server)

	log.Info("Starting grpc server")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Warn("Failed to load .env file")
	}

	app.LoadConfig()
	logging.InitLogging(logging.Config{
		LogFile: app.Config.LogFile,
	})

	db.InitDB()

	var wg sync.WaitGroup

	wg.Add(1)
	go httpStarter(&wg, app.Config.HttpBindAddress)
	wg.Add(1)
	go gRPCStarter(&wg, app.Config.GrpcBindAddress, app.Config.DataManagerAddress)

	wg.Wait()
}
