package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRequest(method string, url string, jsonBody *map[string]string, headers *map[string]string) (int, string) {
	var requestBody io.Reader
	if jsonBody != nil {
		bodyBytes, _ := json.Marshal(jsonBody)
		requestBody = bytes.NewBuffer(bodyBytes)
	}

	request, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		log.Fatal(err)
	}
	if headers != nil {
		for name, value := range *headers {
			request.Header.Set(name, value)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return response.StatusCode, string(body)
}

func MakeOtherRequest(method string, url string, jsonBody *map[string][]string, headers *map[string]string) (int, string) {
	var requestBody io.Reader
	if jsonBody != nil {
		bodyBytes, _ := json.Marshal(jsonBody)
		requestBody = bytes.NewBuffer(bodyBytes)
	}

	request, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		log.Fatal(err)
	}
	if headers != nil {
		for name, value := range *headers {
			request.Header.Set(name, value)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return response.StatusCode, string(body)
}