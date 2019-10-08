package models

import (
	"Golang-Docker/pkg/config"
	"github.com/jinzhu/gorm"
	"github.com/microcosm-cc/bluemonday"
)

var dbSubscriberClient *gorm.DB
var subscribersPerPage = 5

type Subscriber struct {
	gorm.Model
	// Id string `json:"id"`
	FirstName  string `gorm:""json:"first_name"`
	LastName   string `gorm:""json:"last_name"`
	Email      string `gorm:""json:"email"`
	Newsletter bool   `gorm:"default:'false'"`
}

func init() {
	config.Connect()
	dbSubscriberClient = config.GetDB()
	dbSubscriberClient.AutoMigrate(&Subscriber{})
}

func (s *Subscriber) NewSubscriber() *Subscriber {
	policy := bluemonday.UGCPolicy()
	s.FirstName = policy.Sanitize(s.FirstName)
	s.LastName = policy.Sanitize(s.LastName)
	s.Newsletter = false
	if !validEmail(s.Email) {
		return nil
	}
	dbSubscriberClient.NewRecord(s)
	dbSubscriberClient.Create(&s)
	return s
}

func GetAllSubscribers() []Subscriber {
	var subscribers []Subscriber
	dbSubscriberClient.Find(&subscribers)
	return subscribers
}

func GetSubscriberById(id int64) (*Subscriber, *gorm.DB) {
	var subscriber Subscriber
	db := dbSubscriberClient.Where("id = ?", id).Find(&subscriber)
	return &subscriber, db
}

func GetSubscribersByLikeStatement(email string, firstName string, lastName string) ([]Subscriber, *gorm.DB) {
	var subscribers []Subscriber
	var db *gorm.DB
	if email != "" && firstName != "" && lastName != "" {
		db = dbSubscriberClient.Where("email LIKE ?", email).Where("first_name LIKE ?", firstName).Where("last_name LIKE ?", lastName).Find(&subscribers)
	} else if firstName != "" && lastName != "" {
		db = dbSubscriberClient.Where("first_name LIKE ?", firstName).Where("last_name LIKE ?", lastName).Find(&subscribers)
	} else if email != "" && firstName != "" {
		db = dbSubscriberClient.Where("email LIKE ?", email).Where("first_name LIKE ?", firstName).Find(&subscribers)
	} else if email != "" && lastName != "" {
		db = dbSubscriberClient.Where("email LIKE ?", email).Where("last_name LIKE ?", lastName).Find(&subscribers)
	} else if email != "" {
		db = dbSubscriberClient.Where("email LIKE ?", email).Find(&subscribers)
	} else if firstName != "" {
		db = dbSubscriberClient.Where("first_name LIKE ?", firstName).Find(&subscribers)
	} else if lastName != "" {
		db = dbSubscriberClient.Where("last_name LIKE ?", lastName).Find(&subscribers)
	} else {
		return nil, nil
	}
	return subscribers, db
}

func DeleteSubscriber(id int64) Subscriber {
	var subscriber Subscriber
	dbSubscriberClient.Where("id = ?", id).Delete(subscriber)
	return subscriber
}

func GetSubscribersByPageNumber(page int64) []Subscriber {
	var wantedSubscribers []Subscriber
	if page < 1 {
		return nil
	} else {
		offset := (page - 1) * int64(subscribersPerPage)
		dbSubscriberClient.Order("created_at desc").Offset(offset).Limit(subscribersPerPage).Find(&wantedSubscribers)
		return wantedSubscribers
	}
}
