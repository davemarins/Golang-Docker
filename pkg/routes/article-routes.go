package routes

import (
	"github.com/gorilla/mux"
	"simple-REST/pkg/controllers"
)

var RegisterArticleRoutes = func(router *mux.Router) {
	router.HandleFunc("/article/", controllers.NewArticle).Methods("POST")
	router.HandleFunc("/article/all/", controllers.GetAllArticles).Methods("GET")
	router.HandleFunc("/public/article/page/{pageNumber}", controllers.GetArticlesByPageNumber).Methods("GET")
	router.HandleFunc("/public/article/{articleId}", controllers.GetArticle).Methods("GET")
	router.HandleFunc("/article/{articleId}", controllers.UpdateArticle).Methods("PUT")
	router.HandleFunc("/article/{articleId}", controllers.DeleteArticle).Methods("DELETE")
}
