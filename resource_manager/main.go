package main

import (
	pb "github.com/dc-lab/sky/api/proto/resource_manager"
	"github.com/dc-lab/sky/resource_manager/controllers"
	"github.com/dc-lab/sky/resource_manager/grpc_server"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

func httpStarter(wg *sync.WaitGroup, addr string) {
	defer wg.Done()
	router := mux.NewRouter()

	router.HandleFunc("/resources", controllers.ResourcesHandle).Methods("GET", "POST")
	router.HandleFunc("/resources/{id}", controllers.ResourceHandle).Methods("GET", "DELETE")

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
	var wg sync.WaitGroup

	wg.Add(1)
	go httpStarter(&wg, ":8090")
	wg.Add(1)
	go gRPCStarter(&wg, ":5051")

	wg.Wait()
}
