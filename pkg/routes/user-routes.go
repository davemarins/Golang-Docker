package routes

import (
	"Golang-Docker/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.NewUser).Methods("POST")
	router.HandleFunc("/user/all/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/public/user/login/", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/user/refresh/", controllers.RefreshUser).Methods("POST")
}

// TODO recaptcha needed for:
// LoginUser
