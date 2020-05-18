package main

import (
	"github.com/dc-lab/sky/user_manager/app"
	"github.com/dc-lab/sky/user_manager/handles"
	"github.com/dc-lab/sky/user_manager/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path"
)

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

	router := mux.NewRouter()

	router.HandleFunc("/register", handles.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", handles.Login).Methods(http.MethodPost)
	router.HandleFunc("/change_password", handles.ChangePassword).Methods(http.MethodPost)
	router.HandleFunc("/groups", handles.Groups).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/groups/{id}", handles.Group).Methods(http.MethodGet, http.MethodDelete, http.MethodPost)

	log.Fatal(http.ListenAndServe(app.Config.HTTPAddress, router))
}
