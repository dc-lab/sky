package app

import (
	"log"
	"net/http"
)

func Unauthorized(w http.ResponseWriter) {
	http.Error(w, "No user id specified", http.StatusUnauthorized)
}

type ResourceNotFound struct{}

func (e *ResourceNotFound) Error() string {
	return "resource not found"
}

func HandleBaseError(w http.ResponseWriter, err error, status int) {
	log.Println(err)
	http.Error(w, err.Error(), status)
}
