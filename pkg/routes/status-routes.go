package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var RegisterStatusRoutes = func(router *mux.Router) {
	router.HandleFunc("/public/status/", status).Methods("GET")
}
