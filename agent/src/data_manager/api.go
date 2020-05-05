package data_manager_api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dc-lab/sky/agent/src/common"
)

const DATA_MANAGER_API_URL = "https://sky.sskvor.dev/v1/"

func GetFileBody(file_id string) io.ReadCloser {
	url := fmt.Sprintf(DATA_MANAGER_API_URL+"/files/%s/data", file_id)
	resp, err := http.Get(url)
	// TODO: handle data manager errors
	common.DealWithError(err)
	return resp.Body
}

func UploadFile(file_path string) {
}
