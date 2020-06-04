package http_handles

import (
	"encoding/json"
	"github.com/dc-lab/sky/resource_manager/app"
	"github.com/dc-lab/sky/resource_manager/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Resources(w http.ResponseWriter, req *http.Request) {
	userId := req.Header.Get("User-Id")
	if userId == "" {
		log.Println("Empty User-Id header")
		app.Unauthorized(w)
		return
	}

	switch req.Method {
	case http.MethodGet:
		message, resources := db.GetUserResources(userId)
		if resources == nil {
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		resourceJson, err := json.Marshal(resources)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(resourceJson)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled resources get request")
	case http.MethodPost:
		d := json.NewDecoder(req.Body)
		resource := &db.Resource{}
		err := d.Decode(resource)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if message, ok := resource.Create(userId); !ok {
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		resourceJson, err := json.Marshal(resource)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(resourceJson)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled resources post request")
	}
}

type changeBody struct {
	UsersToAdd     []string `json:"users_to_add"`
	UsersToRemove  []string `json:"users_to_remove"`
	GroupsToAdd    []string `json:"groups_to_add"`
	GroupsToRemove []string `json:"groups_to_remove"`
}

func Resource(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	userId := req.Header.Get("User-Id")
	if userId == "" {
		log.Println("Empty User-Id header")
		app.Unauthorized(w)
		return
	}
	switch req.Method {
	case http.MethodGet:
		message, resource := db.GetResource(userId, id)
		if resource == nil {
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}
		resourceJson, err := json.Marshal(*resource)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(resourceJson); err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled resource get request")
	case http.MethodDelete:
		if err := db.DeleteResource(userId, id); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write([]byte("Success")); err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled resource delete request")
	case http.MethodPost:
		message, resource := db.GetResource(userId, id)
		if resource == nil {
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(req.Body)
		changes := &changeBody{}
		err := decoder.Decode(changes)
		if err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}
		if err := resource.Modify(changes.UsersToAdd, changes.UsersToRemove, changes.GroupsToAdd, changes.GroupsToRemove); err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}
		message, resource = db.GetResource(userId, id)
		if resource == nil {
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		resourceJson, err := json.Marshal(*resource)
		if err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(resourceJson); err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled resource change request")
	}
}
