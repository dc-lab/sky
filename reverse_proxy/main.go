package main

import (
	"fmt"
	"github.com/dc-lab/sky/reverse_proxy/app"
	"github.com/dc-lab/sky/reverse_proxy/db"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
)

func handlerProxy(w http.ResponseWriter, r *http.Request) {
	var host string
	var authOptional bool
	for _, endpoint := range app.Config.Endpoints {
		if endpoint.PathRegex.MatchString(r.URL.String()) {
			host = endpoint.Hostname
			authOptional = endpoint.AuthOptional
			break
		}
	}
	if host == "" {
		fmt.Printf("Can't find endpoint for url '%s'\n", r.URL.String())
		http.Error(w, "No such url", http.StatusNotFound)
		return
	}

	if !authOptional {
		userToken := r.Header.Get("User-Token")
		if userToken == "" {
			fmt.Println("No credentials provided")
			http.Error(w, "No credentials provided", http.StatusUnauthorized)
			return
		}
		userId, err := db.GetIdByToken(userToken)
		if err != nil {
			fmt.Println(err)
			switch err.(type) {
			case *db.UserNotFoundError:
				http.Error(w, "Authorization failed", http.StatusForbidden)
			default:
				http.Error(w, "Internal error", http.StatusInternalServerError)
			}
			return
		}
		r.Header.Set("User-Id", userId)
	}

	newUrl, err := url.Parse(fmt.Sprintf("http://%s/", host))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something is wrong", http.StatusInternalServerError)
	}
	proxy := httputil.NewSingleHostReverseProxy(newUrl)
	proxy.ServeHTTP(w, r)
}

func main() {
	app.ParseConfig()

	logPath := path.Join(app.Config.LogsDir, "rp.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

	db.InitDB()

	http.HandleFunc("/", handlerProxy)
	if err := http.ListenAndServe(app.Config.HTTPAddress, nil); err != nil {
		panic(err)
	}
}
