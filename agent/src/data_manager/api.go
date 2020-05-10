package data_manager_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/dc-lab/sky/agent/src/protos"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dc-lab/sky/agent/src/common"
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

func UploadFile(absolutePath string, fileDir string) pb.TFile {
	relativePath := common.ConvertToRelativePath(fileDir, absolutePath)
	metadata := UploadFileMetadata(relativePath)
	fileId := metadata["id"]
	UploadFileContent(absolutePath, "https://"+metadata["upload_url"])
	return pb.TFile{Id: &fileId, AgentRelativeLocalPath: &relativePath}
}

func UploadFileMetadata(filePath string) map[string]string {
	values := map[string]string{"name": filePath}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(DATA_MANAGER_API_URL+"/files", "application/json", bytes.NewBuffer(jsonValue))
	common.DealWithError(err)
	body, err := ioutil.ReadAll(resp.Body)
	common.DealWithError(err)
	var data map[string]string
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
