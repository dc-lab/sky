package main

import (
	"github.com/dc-lab/sky/user_manager/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.RegisterHandle).Methods("POST")
	router.HandleFunc("/login", controllers.LoginHandle).Methods("POST")

	log.Fatal(http.ListenAndServe(":8090", router))
}
