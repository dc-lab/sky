package main

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type TaskSpec struct {
	Command     string
	InputFiles  []string
	OutputFiles []string
	TimeLimit   uint64
}

type JobSpec struct {
	Tasks []TaskSpec
	Type  string
}

type JobState struct {
	JobId   uuid.UUID
	State   string
	Results []uuid.UUID
	Spec    JobSpec
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

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

func startJob(w http.ResponseWriter, r *http.Request) {
	jobId, err := uuid.NewV4()
	panicOnError(err)

	jobSpec := new(JobSpec)
	err = json.NewDecoder(r.Body).Decode(jobSpec)
	panicOnError(err)

	jobState := new(JobState)
	jobState.JobId = jobId
	jobState.State = "Starting"
	jobState.Results = make([]uuid.UUID, 0)
	jobState.Spec = *jobSpec

	_ = json.NewEncoder(os.Stdout).Encode(jobState)

	err = json.NewEncoder(w).Encode(jobId)
	panicOnError(err)
}

func getJobs(w http.ResponseWriter, _ *http.Request) {
	jobStates := make([]JobState, 0)
	i := 0
	for i < 5 {
		jobId, _ := uuid.NewV4()
		jobState := JobState{JobId: jobId}
		jobStates = append(jobStates, jobState)
		i++
	}
	_ = json.NewEncoder(w).Encode(jobStates)
}

func getJob(w http.ResponseWriter, r *http.Request) {
	jobId, ok := mux.Vars(r)["job_id"]
	if !ok {
		http.Error(w, "job_id not found in URL", http.StatusBadRequest)
	}
	jobState := JobState{JobId: uuid.FromStringOrNil(jobId)}
	_ = json.NewEncoder(w).Encode(jobState)
}

func cancelJob(w http.ResponseWriter, r *http.Request) {
	jobId, ok := mux.Vars(r)["job_id"]
	if !ok {
		http.Error(w, "job_id not found in URL", http.StatusBadRequest)
	}
	log.Printf("cancel job with id = %s", jobId)
	_ = json.NewEncoder(w).Encode(jobId)
}

func deleteJob(w http.ResponseWriter, r *http.Request) {
	jobId, ok := mux.Vars(r)["job_id"]
	if !ok {
		http.Error(w, "job_id not found in URL", http.StatusBadRequest)
	}
	log.Printf("delete job with id = %s", jobId)
	_ = json.NewEncoder(w).Encode(jobId)
}
