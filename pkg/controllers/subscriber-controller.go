package controllers

import (
	"Golang-Docker/pkg/models"
	"Golang-Docker/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func NewSubscriber(w http.ResponseWriter, r *http.Request) {
	newSubscriber := &models.Subscriber{}
	utils.ParseBody(r, newSubscriber)
	s := newSubscriber.NewSubscriber()
	if s == nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		result, _ := json.Marshal(s)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func GetSubscriber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subscriberId := vars["id"]
	id, err := strconv.ParseInt(subscriberId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	subscriberDetails, _ := models.GetSubscriberById(id)
	if subscriberDetails.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result, _ := json.Marshal(subscriberDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func EditSubscriber(w http.ResponseWriter, r *http.Request) {
	var updateSubscriber = &models.Subscriber{}
	utils.ParseBody(r, updateSubscriber)
	vars := mux.Vars(r)
	subscriberId := vars["id"]
	id, err := strconv.ParseInt(subscriberId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	subscriberDetails, db := models.GetSubscriberById(id)
	if subscriberDetails.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if updateSubscriber.Email != "" {
		subscriberDetails.Email = updateSubscriber.Email
	}
	if updateSubscriber.FirstName != "" {
		subscriberDetails.FirstName = updateSubscriber.FirstName
	}
	if updateSubscriber.LastName != "" {
		subscriberDetails.LastName = updateSubscriber.LastName
	}
	if updateSubscriber.Newsletter != subscriberDetails.Newsletter {
		subscriberDetails.Newsletter = updateSubscriber.Newsletter
	}
	db.Save(&subscriberDetails)
	result, _ := json.Marshal(subscriberDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteSubscriber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subscriberId := vars["id"]
	id, err := strconv.ParseInt(subscriberId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	subscriberDetails, _ := models.GetSubscriberById(id)
	if subscriberDetails.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	subscriber := models.DeleteSubscriber(int64(subscriberDetails.ID))
	result, _ := json.Marshal(subscriber)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func FindSubscriber(w http.ResponseWriter, r *http.Request) {
	givenSubscriber := &models.Subscriber{}
	utils.ParseBody(r, givenSubscriber)
	foundSubscribers, _ := models.GetSubscribersByLikeStatement(givenSubscriber.Email, givenSubscriber.FirstName, givenSubscriber.LastName)
	result, _ := json.Marshal(foundSubscribers)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func AllSubscribers(w http.ResponseWriter, r *http.Request) {
	subscribers := models.GetAllSubscribers()
	result, _ := json.Marshal(subscribers)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetSubscribersByPageNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageId := vars["pageNumber"]
	pageNumber, err := strconv.ParseInt(pageId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	wantedSubscribers := models.GetSubscribersByPageNumber(pageNumber)
	result, _ := json.Marshal(wantedSubscribers)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UnsubscribeFromNewsletter(w http.ResponseWriter, r *http.Request) {
	givenSubscriber := &models.Subscriber{}
	utils.ParseBody(r, givenSubscriber)
	foundSubscriber, db := models.GetSubscribersByLikeStatement(givenSubscriber.Email, givenSubscriber.FirstName, givenSubscriber.LastName)
	if len(foundSubscriber) > 1 {
		w.WriteHeader(http.StatusConflict)
		return
	}
	foundSubscriber[0].Newsletter = false
	db.Save(&foundSubscriber)
	result, _ := json.Marshal(foundSubscriber)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
