package data_manager_api

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/dc-lab/sky/agent/src/common"
)

const DATA_MANAGER_API_URL = "https://sky.sskvor.dev/v1/"

func GetFileBody(file_id string) (error, io.ReadCloser) {
	url := fmt.Sprintf(DATA_MANAGER_API_URL+"/files/%s/data", file_id)
	resp, err := http.Get(url)
	if resp.StatusCode != 200 {
		error_text, _ := ioutil.ReadAll(resp.Body)
		err = errors.New(string(error_text))
	}
	common.DealWithError(err)
	return err, resp.Body
}

func UploadFile(file_path string) {
}
