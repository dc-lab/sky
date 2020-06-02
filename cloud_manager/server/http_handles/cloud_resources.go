package http_handles

import (
	"encoding/json"
	"log"
	"net/http"
)

func Resources(w http.ResponseWriter, req *http.Request, dp *DataProvider) {
	userId := req.Header.Get("User-Id")
	if userId == "" {
		log.Println("Empty User-Id header")
		UnauthorizedError(w)
		return
	}

	switch req.Method {
	case http.MethodGet:
		userResources, err := dp.cloudResourcesManager.GetUserResources(userId)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userResourcesJson, err := json.Marshal(userResources)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(userResourcesJson)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Successfully handled resources get request")
	}
}
