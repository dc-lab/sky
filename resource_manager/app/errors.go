package app

import "net/http"

func Unauthorized(w http.ResponseWriter) {
	http.Error(w, "No user id specified", http.StatusUnauthorized)
}
