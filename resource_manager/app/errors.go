package app

import "net/http"

func Unauthorized(w http.ResponseWriter) {
	http.Error(w, "No user id specified", http.StatusUnauthorized)
}

type ResourceNotFound struct{}

func (e *ResourceNotFound) Error() string {
	return "resource not found"
}
