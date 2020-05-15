package main

import (
	"github.com/dc-lab/sky/job_manager/controllers"
	"github.com/dc-lab/sky/job_manager/util"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/jobs", controllers.StartJob).Methods(http.MethodPost)
	router.HandleFunc("/jobs/{job_id}", getJob).Methods(http.MethodGet)
	router.HandleFunc("/jobs", getJobs).Methods(http.MethodGet)
	router.HandleFunc("/jobs/{job_id}/cancel", cancelJob).Methods(http.MethodPost)
	router.HandleFunc("/jobs/{job_id}", deleteJob).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
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

	util.panicOnError(util.encodeBody(w, jobStates), w, http.StatusInternalServerError)
}

func getJob(w http.ResponseWriter, r *http.Request) {
	jobId := mux.Vars(r)["job_id"]

	jobState := JobState{JobId: uuid.FromStringOrNil(jobId)}

	util.panicOnError(util.encodeBody(w, jobState), w, http.StatusInternalServerError)
}

func cancelJob(w http.ResponseWriter, r *http.Request) {
	jobId := mux.Vars(r)["job_id"]

	log.Printf("cancel job with id = %s", jobId)

	util.panicOnError(util.encodeBody(w, jobId), w, http.StatusInternalServerError)
}

func deleteJob(w http.ResponseWriter, r *http.Request) {
	jobId := mux.Vars(r)["job_id"]

	log.Printf("cancel job with id = %s", jobId)

	util.panicOnError(util.encodeBody(w, jobId), w, http.StatusInternalServerError)
}
