package http_handles

import (
	"net/http"
)

func UnauthorizedError(w http.ResponseWriter) {
	http.Error(w, "No user id specified", http.StatusUnauthorized)
}
