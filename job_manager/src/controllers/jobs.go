package controllers

import (
	"github.com/dc-lab/sky/job_manager/util"
	"github.com/gofrs/uuid"
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
	util.PanicOnError(err, w, http.StatusInternalServerError)

	jobSpec := new(JobSpec)
	util.PanicOnError(util.DecodeBody(r.Body, jobSpec), w, http.StatusBadRequest)

	jobState := new(JobState)
	jobState.JobId = jobId
	jobState.State = "Starting"
	jobState.Results = make([]uuid.UUID, 0)
	jobState.Spec = *jobSpec

	util.PanicOnError(util.EncodeBody(w, jobState.JobId), w, http.StatusInternalServerError)
}
