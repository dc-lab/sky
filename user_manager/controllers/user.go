package controllers

import (
	"encoding/json"
	"github.com/dc-lab/sky/user_manager/db"
	"log"
	"net/http"
)

func RegisterHandle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		user := &db.User{}
		err := decoder.Decode(user)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if message, ok := user.Create(); !ok {
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		userJson, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(userJson)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled register post request")
	}
}

func LoginHandle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		user := &db.User{}
		err := decoder.Decode(user)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if message, ok := user.Get(); !ok {
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		userJson, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(userJson)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled login post request")
	}
}
