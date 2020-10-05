package handles

import (
	"encoding/json"
	"github.com/dc-lab/sky/internal/user_manager/app"
	"github.com/dc-lab/sky/internal/user_manager/db"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Groups(w http.ResponseWriter, req *http.Request) {
	userId := req.Header.Get("User-Id")
	if userId == "" {
		log.Println("Empty User-Id header")
		app.Unauthorized(w)
		return
	}
	switch req.Method {
	case http.MethodGet:
		groups, err := db.GetGroups(userId)
		if err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}

		groupsJson, err := json.Marshal(groups)
		if err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(groupsJson)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled groups get request")
	case http.MethodPost:
		d := json.NewDecoder(req.Body)
		group := &db.Group{}
		err := d.Decode(group)
		if err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}

		if err := group.Create(userId); err != nil {
			log.Println(err)
			switch err.(type) {
			case *app.EmptyField:
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, "Internal error", http.StatusInternalServerError)
			}
			return
		}

		groupJson, err := json.Marshal(group)
		if err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(groupJson)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled groups post request")
	}
}

type changeBody struct {
	ToAdd    []string `json:"users_to_add"`
	ToRemove []string `json:"users_to_remove"`
}

func Group(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	userId := req.Header.Get("User-Id")
	if userId == "" {
		log.Println("Empty User-Id header")
		app.Unauthorized(w)
		return
	}
	switch req.Method {
	case http.MethodGet:
		group, err := db.GetGroup(userId, id)
		if err != nil {
			log.Println(err)
			switch err.(type) {
			case *app.GroupNotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
			case *app.PermissionDenied:
				http.Error(w, err.Error(), http.StatusForbidden)
			default:
				http.Error(w, "Internal error", http.StatusInternalServerError)
			}
			return
		}
		groupJson, err := json.Marshal(*group)
		if err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(groupJson); err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled group get request")
	case http.MethodDelete:
		group, err := db.GetGroup(userId, id)
		if err != nil {
			log.Println(err)
			switch err.(type) {
			case *app.GroupNotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
			case *app.PermissionDenied:
				http.Error(w, err.Error(), http.StatusForbidden)
			default:
				http.Error(w, "Internal error", http.StatusInternalServerError)
			}
			return
		}
		err = group.Delete()
		if err != nil {
			log.Println(err)
			switch err.(type) {
			case *app.GroupNotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, "Internal error", http.StatusInternalServerError)
			}
			return
		}
		if _, err := w.Write([]byte("Success")); err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled group delete request")
	case http.MethodPost:
		group, err := db.GetGroup(userId, id)
		if err != nil {
			log.Println(err)
			switch err.(type) {
			case *app.GroupNotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
			case *app.PermissionDenied:
				http.Error(w, err.Error(), http.StatusForbidden)
			default:
				http.Error(w, "Internal error", http.StatusInternalServerError)
			}
			return
		}

		decoder := json.NewDecoder(req.Body)
		changes := &changeBody{}
		err = decoder.Decode(changes)
		if err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}
		if err := group.Modify(changes.ToAdd, changes.ToRemove); err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}
		group, err = db.GetGroup(userId, id)
		if err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
		}

		groupJson, err := json.Marshal(*group)
		if err != nil {
			app.HandleBaseError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(groupJson); err != nil {
			log.Println(err)
			return
		}
		log.Println("Successfully handled group change request")
	}
}
