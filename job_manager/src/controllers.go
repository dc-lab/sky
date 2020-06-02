package main

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TaskSpec struct {
	Command     string
	InputFiles  []string `json:"input_files"`
	OutputFiles []string `json:"output_files"`
	TimeLimit   uint64   `json:"time_limit"`
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

func startJob(w http.ResponseWriter, r *http.Request) {
	jobId, err := uuid.NewV4()
	panicOnError(err, w, http.StatusInternalServerError)

	jobSpec := new(JobSpec)
	panicOnError(decodeBody(r.Body, jobSpec), w, http.StatusBadRequest)

	jobState := new(JobState)
	jobState.JobId = jobId
	jobState.State = "Starting"
	jobState.Results = make([]uuid.UUID, 0)
	jobState.Spec = *jobSpec

	panicOnError(encodeBody(w, jobState.JobId), w, http.StatusInternalServerError)
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

	panicOnError(encodeBody(w, jobStates), w, http.StatusInternalServerError)
}

func getJob(w http.ResponseWriter, r *http.Request) {
	jobId := mux.Vars(r)["job_id"]

	jobState := JobState{JobId: uuid.FromStringOrNil(jobId)}

	panicOnError(encodeBody(w, jobState), w, http.StatusInternalServerError)
}

func cancelJob(w http.ResponseWriter, r *http.Request) {
	jobId := mux.Vars(r)["job_id"]

	log.Printf("cancel job with id = %s", jobId)

	panicOnError(encodeBody(w, jobId), w, http.StatusInternalServerError)
}

func deleteJob(w http.ResponseWriter, r *http.Request) {
	jobId := mux.Vars(r)["job_id"]

	log.Printf("cancel job with id = %s", jobId)

	panicOnError(encodeBody(w, jobId), w, http.StatusInternalServerError)
}
