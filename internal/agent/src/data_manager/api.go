package data_manager_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dc-lab/sky/internal/data_manager/modelapi"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dc-lab/sky/internal/agent/src/common"
	rm "github.com/dc-lab/sky/api/proto"
)

const DATA_MANAGER_API_URL = "https://sky.sskvor.dev/v1"

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

func UploadFile(absolutePath string, fileDir string) rm.TFile {
	relativePath := common.ConvertToRelativePath(fileDir, absolutePath)
	metadata := UploadFileMetadata(relativePath)
	fileId := metadata.Id
	UploadFileContent(absolutePath, metadata.UploadUrl)
	return rm.TFile{Id: fileId, AgentRelativeLocalPath: relativePath}
}

func UploadFileMetadata(filePath string) modelapi.FileResponse {
	values := map[string]string{"name": filePath}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(DATA_MANAGER_API_URL+"/files", "application/json", bytes.NewBuffer(jsonValue))
	common.DealWithError(err)
	body, err := ioutil.ReadAll(resp.Body)
	common.DealWithError(err)
	var data modelapi.FileResponse
	err = json.Unmarshal(body, &data)
	common.DealWithError(err)
	return data
}

func UploadFileContent(filePath string, uploadUrl string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", uploadUrl, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(r)
	common.DealWithError(err)
	fmt.Println(resp)
}
