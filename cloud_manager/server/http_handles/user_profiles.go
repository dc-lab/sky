package http_handles

import (
	"encoding/json"
	"github.com/dc-lab/sky/cloud_manager/server/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UserProfiles(w http.ResponseWriter, req *http.Request, dp *DataProvider) {
	userId := req.Header.Get("User-Id")
	if userId == "" {
		log.Println("Empty User-Id header")
		UnauthorizedError(w)
		return
	}

	switch req.Method {
	case http.MethodGet:
		// Get all creds available for user
		userCreds, err := dp.credsManager.GetAllUserCreds(userId)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userCredsJson, err := json.Marshal(userCreds)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(userCredsJson)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Successfully handled user profiles get request")
	case http.MethodPost:
		// Create new creds for user
		d := json.NewDecoder(req.Body)
		credsInput := &entity.CreateCredentialsInput{}
		if err := d.Decode(credsInput); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := validateCredentialsInput(credsInput); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		creds, err := dp.credsManager.CreateCreds(userId, credsInput)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		credsJson, err := json.Marshal(creds)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(credsJson)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Successfully handled user profiles post request")
	}
}

func UserProfile(w http.ResponseWriter, req *http.Request, dp *DataProvider) {
	credsId := mux.Vars(req)["id"]
	userId := req.Header.Get("User-Id")
	if userId == "" {
		log.Println("Empty User-Id header")
		UnauthorizedError(w)
		return
	}

	switch req.Method {
	case http.MethodGet:
		// Get user creds by id
		userCreds, err := dp.credsManager.GetUserCreds(userId, credsId)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userCredsJson, err := json.Marshal(userCreds)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(userCredsJson)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Successfully handled user profile get request")
	case http.MethodDelete:
		// Delete user creds by id
		err := dp.credsManager.DeleteUserCreds(userId, credsId)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		_, err = w.Write([]byte("OK"))
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Successfully handled user profile delete request")
	}
}
