package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"simple-REST/pkg/models"
	"simple-REST/pkg/utils"
	"strconv"
)

func NewArticle(w http.ResponseWriter, r *http.Request) {
	newArticle := &models.Article{}
	utils.ParseBody(r, newArticle)
	a := newArticle.NewArticle()
	result, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	allArticles := models.GetAllArticles()
	result, _ := json.Marshal(allArticles)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetArticlesByPageNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageId := vars["pageNumber"]
	pageNumber, err := strconv.ParseInt(pageId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	wantedArticles := models.GetArticlesByPageNumber(pageNumber)
	result, _ := json.Marshal(wantedArticles)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	id, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	article, _ := models.GetArticleById(id)
	if article.Title == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result, _ := json.Marshal(article)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var updateArticle = &models.Article{}
	utils.ParseBody(r, updateArticle)
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	id, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	articleDetails, db := models.GetArticleById(id)
	if articleDetails.Title == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if updateArticle.AuthorId != updateArticle.AuthorId {
		articleDetails.AuthorId = updateArticle.AuthorId
	}
	if updateArticle.Content != "" {
		articleDetails.Content = updateArticle.Content
	}
	if updateArticle.Description != "" {
		articleDetails.Description = updateArticle.Description
	}
	if updateArticle.Title != "" {
		articleDetails.Title = updateArticle.Title
	}
	db.Save(&articleDetails)
	result, _ := json.Marshal(articleDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	id, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	findArticle, _ := models.GetArticleById(id)
	if findArticle.Title == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	article := models.DeleteArticleById(id)
	result, _ := json.Marshal(article)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
