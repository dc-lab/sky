package handles

import (
	"encoding/json"
	"github.com/dc-lab/sky/user_manager/app"
	"github.com/dc-lab/sky/user_manager/db"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, req *http.Request) {
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

func Login(w http.ResponseWriter, req *http.Request) {
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

func ChangePassword(w http.ResponseWriter, req *http.Request) {
	userId := req.Header.Get("User-Id")
	if userId == "" {
		log.Println("Empty User-Id header")
		app.Unauthorized(w)
		return
	}
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var body map[string]interface{}
		err := decoder.Decode(&body)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		oldPassword := body["old_password"].(string)
		newPassword := body["new_password"].(string)
		err = db.ChangePassword(userId, oldPassword, newPassword)
		if err != nil {
			log.Printf("Error occured: %s\n", err)
			switch err.(type) {
			case *app.UserNotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
			case *app.WrongPassword:
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, "Internal error", http.StatusInternalServerError)
			}
			return
		}
		_, err = w.Write([]byte("Ok, password changed"))
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Change password handled successfully")
	}
}