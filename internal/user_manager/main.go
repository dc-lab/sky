package main

import (
	pb "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/user_manager/app"
	"github.com/dc-lab/sky/internal/user_manager/db"
	"github.com/dc-lab/sky/internal/user_manager/grpc_server"
	"github.com/dc-lab/sky/internal/user_manager/handles"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"path"
	"sync"
)

func httpStarter(wg *sync.WaitGroup, addr string) {
	defer wg.Done()
	router := mux.NewRouter()

	router.HandleFunc("/register", handles.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", handles.Login).Methods(http.MethodPost)
	router.HandleFunc("/change_password", handles.ChangePassword).Methods(http.MethodPost)
	router.HandleFunc("/groups", handles.Groups).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/groups/{id}", handles.Group).Methods(http.MethodGet, http.MethodDelete, http.MethodPost)

	log.Fatal(http.ListenAndServe(addr, router))
}

func gRPCStarter(wg *sync.WaitGroup, addr string) {
	defer wg.Done()
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserManagerServer(server, grpc_server.Server{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

func main() {
	app.ParseConfig()

	logPath := path.Join(app.Config.LogsDir, "um.log")
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
