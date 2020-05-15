package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/jobs", startJob).Methods(http.MethodPost)
	router.HandleFunc("/jobs/{job_id}", getJob).Methods(http.MethodGet)
	router.HandleFunc("/jobs", getJobs).Methods(http.MethodGet)
	router.HandleFunc("/jobs/{job_id}/cancel", cancelJob).Methods(http.MethodPost)
	router.HandleFunc("/jobs/{job_id}", deleteJob).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
}
