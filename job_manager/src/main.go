package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/jobs", startJob).Methods(http.MethodPost)
	router.HandleFunc("/jobs/{job_id}", getJob).Methods(http.MethodGet)
	router.HandleFunc("/jobs", getJobs).Methods(http.MethodGet)
	router.HandleFunc("/jobs/{job_id}/cancel", cancelJob).Methods(http.MethodPost)
	router.HandleFunc("/jobs/{job_id}", deleteJob).Methods(http.MethodDelete)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func startJob(_ http.ResponseWriter, _ *http.Request) {
	log.Printf("job started")
}

func getJobs(_ http.ResponseWriter, _ *http.Request) {
	log.Printf("get jobs")
}

func getJob(w http.ResponseWriter, r *http.Request) {
	jobId, ok := mux.Vars(r)["job_id"]
	if !ok {
		http.Error(w, "job_id not found in URL", http.StatusBadRequest)
	}
	log.Printf("get job with id = %s", jobId)
}

func cancelJob(w http.ResponseWriter, r *http.Request) {
	jobId, ok := mux.Vars(r)["job_id"]
	if !ok {
		http.Error(w, "job_id not found in URL", http.StatusBadRequest)
	}
	log.Printf("cancel job with id = %s", jobId)
}

func deleteJob(w http.ResponseWriter, r *http.Request) {
	jobId, ok := mux.Vars(r)["job_id"]
	if !ok {
		http.Error(w, "job_id not found in URL", http.StatusBadRequest)
	}
	log.Printf("delete job with id = %s", jobId)
}
