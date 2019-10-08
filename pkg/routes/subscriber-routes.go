package routes

import (
	"Golang-Docker/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterSubscribersRoutes = func(router *mux.Router) {
	router.HandleFunc("/public/subscriber/", controllers.NewSubscriber).Methods("POST")
	router.HandleFunc("/subscriber/{id}", controllers.GetSubscriber).Methods("GET")
	router.HandleFunc("/subscriber/{id}", controllers.EditSubscriber).Methods("PUT")
	router.HandleFunc("/subscriber/{id}", controllers.DeleteSubscriber).Methods("DELETE")
	router.HandleFunc("/subscriber/find/", controllers.FindSubscriber).Methods("POST")
	router.HandleFunc("/subscriber/all/", controllers.AllSubscribers).Methods("GET")
	router.HandleFunc("/subscriber/page/{pageNumber}", controllers.GetSubscribersByPageNumber).Methods("GET")
	router.HandleFunc("/public/subscriber/unsubscribe/", controllers.UnsubscribeFromNewsletter).Methods("POST")
}
