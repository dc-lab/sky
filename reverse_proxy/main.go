package main

import (
	"encoding/json"
	"fmt"
	"github.com/dc-lab/sky/user_manager/db"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type Endpoint struct {
	PathPrefix string `json:"pathPrefix"`
	Hostname   string `json:"hostname"`
}

type Endpoints struct {
	Endpoints []Endpoint `json:"endpoints"`
}

var endpoints *Endpoints

func handlerProxy(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Host)

	userToken := db.GetIdByToken()
	// Check authorization
	// And set user-id to special header
	// help: Use r.Header.Set(grafanaHeader, grafanaUser)

	var host string
	for _, endpoint := range endpoints.Endpoints {
		if strings.HasPrefix(r.URL.String(), endpoint.PathPrefix) {
			host = endpoint.Hostname
			break
		}
	}
	if host == "" {
		fmt.Printf("Can't find endpoint for url '%s'\n", r.URL.String())
		http.Error(w, "No such url", http.StatusNotFound)
		return
	}

	newUrl, err := url.Parse(fmt.Sprintf("http://%s/", host))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something is wrong", http.StatusInternalServerError)
	}
	proxy := httputil.NewSingleHostReverseProxy(newUrl)
	proxy.ServeHTTP(w, r)
}

func readEndpointConfig(filePath string) (*Endpoints, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rawJson, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var endpoints Endpoints
	err = json.Unmarshal(rawJson, &endpoints)
	if err != nil {
		return nil, err
	}
	return &endpoints, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected path to JSON file with endpoint configuration")
		os.Exit(1)
	}
	var err error
	endpoints, err = readEndpointConfig(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading from %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}

	http.HandleFunc("/", handlerProxy)
	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}
