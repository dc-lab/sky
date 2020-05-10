package main

import (
	"github.com/dc-lab/sky/resource_manager/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/resources", controllers.ResourcesHandle).Methods("GET", "POST")
	router.HandleFunc("/resources/{id}", controllers.ResourceHandle).Methods("GET", "DELETE")

	log.Fatal(http.ListenAndServe(":8090", router))
}
