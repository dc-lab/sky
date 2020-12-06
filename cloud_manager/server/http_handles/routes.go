package http_handles

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(sc *DataProvider) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user-profiles", sc.BindHandler(UserProfiles)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/user-profiles/{id}", sc.BindHandler(UserProfile)).Methods(http.MethodGet, http.MethodDelete)
	router.HandleFunc("/resource-factories", sc.BindHandler(ResourceFactories)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/resource-factories/{id}", sc.BindHandler(ResourceFactory)).Methods(http.MethodGet, http.MethodDelete, http.MethodPost)
	router.HandleFunc("/resources", sc.BindHandler(Resources)).Methods(http.MethodGet)
	return router
}
