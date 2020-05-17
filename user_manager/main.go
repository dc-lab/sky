package main

import (
	"github.com/dc-lab/sky/user_manager/app"
	"github.com/dc-lab/sky/user_manager/controllers"
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

	router.HandleFunc("/register", controllers.RegisterHandle).Methods("POST")
	router.HandleFunc("/login", controllers.LoginHandle).Methods("POST")

	log.Fatal(http.ListenAndServe(app.Config.HTTPAddress, router))
}
