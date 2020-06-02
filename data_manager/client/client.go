package data_manager

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"path"

	pb "github.com/dc-lab/sky/api/proto/data_manager"
)

type File pb.File

type Client struct {
	masterAddress string
	userId        string
}

func MakeClient(masterAddress string, userId string) (*Client, error) {
	return &Client{
		masterAddress: masterAddress,
		userId:        userId,
	}, nil
}

func (c *Client) makeRoute(route string) string {
	return path.Join(c.masterAddress, route)
}

func (c *Client) makeRequest(method string, route string, body io.Reader) (*http.Request, error) {
	url := c.makeRoute(route)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Id", c.userId)

	return req, nil
}

func (c *Client) CreateFile(file *File) (*File, error) {
	encoded, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}

	req, err := c.makeRequest("POST", "/v1/files", bytes.NewReader(encoded))
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resultFile := &File{}
	err = json.Unmarshal(body, &resultFile)
	return resultFile, err
}

func (c *Client) CreateFileWithContents(file *File, body io.Reader) (*File, error) {
	newFile, err := c.CreateFile(file)
	if err != nil {
		return nil, err
	}

	err = c.UploadFileContents(newFile, body)
	if err != nil {
		return nil, err
	}
	newFile.UploadUrls = nil

	return newFile, err
}

type coutingReadWriter struct {
	Count int64
}

func (c *coutingReadWriter) Write(buf []byte) (int, error) {
	c.Count += int64(len(buf))
	return len(buf), nil
}

func (c *Client) UploadFileContents(file *File, body io.Reader) error {
	var lastError error = nil
	counter := &coutingReadWriter{0}
	reader := io.TeeReader(body, counter)
	for _, url := range file.UploadUrls {
		lastError = c.tryUploadFileContents(url, reader)
		if lastError != nil || counter.Count != 0 {
			return lastError
		}
	}

	return lastError
}

func (c *Client) tryUploadFileContents(url string, body io.Reader) error {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Add("User-Id", c.userId)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if !(res.StatusCode >= 200 && res.StatusCode < 400) {
		return errors.New("File upload failed: " + res.Status)
	}

	return nil
}

func (c *Client) GetFile(id string) (*File, error) {
	req, err := c.makeRequest("GET", "/v1/files/"+id, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resultFile := &File{}
	err = json.Unmarshal(body, &resultFile)
	return resultFile, err
}

func (c *Client) GetFileContents(id string, writer io.Writer) error {
	req, err := c.makeRequest("GET", "/v1/files/"+id+"/location", nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var urls []string = nil
	err = json.Unmarshal(body, &urls)
	if err != nil {
		return err
	}

	var lastError error = nil
	counter := &coutingReadWriter{0}
	w := io.MultiWriter(writer, counter)
	for _, url := range urls {
		lastError = c.tryDownloadFileContents(url, w)
		if lastError != nil || counter.Count != 0 {
			return lastError
		}
	}

	return lastError
}

func (c *Client) tryDownloadFileContents(url string, writer io.Writer) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("User-Id", c.userId)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if !(res.StatusCode >= 200 && res.StatusCode < 400) {
		return errors.New("File upload failed: " + res.Status)
	}

	_, err = io.Copy(writer, res.Body)
	return err
}
