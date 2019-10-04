package routes

import (
	"github.com/gorilla/mux"
	"simple-REST/pkg/controllers"
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
