package main

import (
	pb "github.com/dc-lab/sky/api/proto/resource_manager"
	"github.com/dc-lab/sky/resource_manager/app"
	"github.com/dc-lab/sky/resource_manager/db"
	"github.com/dc-lab/sky/resource_manager/grpc_server"
	"github.com/dc-lab/sky/resource_manager/http_handles"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"sync"
)

func httpStarter(wg *sync.WaitGroup, addr string) {
	defer wg.Done()
	router := mux.NewRouter()

	router.HandleFunc("/resources", http_handles.Resources).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/resources/{id}", http_handles.Resource).Methods(http.MethodGet, http.MethodDelete, http.MethodPost)

	log.Fatal(http.ListenAndServe(addr, router))
}

func gRPCStarter(wg *sync.WaitGroup, addr string) {
	defer wg.Done()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterResourceManagerServer(s, grpc_server.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	app.ParseConfig()

	logPath := path.Join(app.Config.LogsDir, "rm.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

	db.InitDB()

	var wg sync.WaitGroup

	wg.Add(1)
	go httpStarter(&wg, app.Config.HTTPAddress)
	wg.Add(1)
	go gRPCStarter(&wg, app.Config.GRPCAddress)

	wg.Wait()
}
