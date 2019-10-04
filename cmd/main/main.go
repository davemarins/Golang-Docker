package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
	"log"
	"net/http"
	"simple-REST/pkg/auth"
	"simple-REST/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	r.Use(auth.JwtAuthentication)
	routes.RegisterUserRoutes(r)
	routes.RegisterArticleRoutes(r)
	http.Handle("/", r)
	handler := cors.Default().Handler(r)
	fmt.Println("Starting web server...")
	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
