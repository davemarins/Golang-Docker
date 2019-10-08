package models

import (
	"net/url"
	"regexp"
	"time"
)

type LoginUser struct {
	Email    string `gorm:""json:"email"`
	Password string `gorm:""json:"password"`
	Code     string `gorm:""json:"totp_code"`
}

type DateMode string

const (
	dash  DateMode = "yyyy-mm-dd"
	slash DateMode = "yyyy/mm/dd"
)

type Role string

const (
	ADMIN  Role = "Admin"
	EDITOR Role = "Editor"
)

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func validEmail(email string) bool {
	return rxEmail.MatchString(email)
}

func validURL(givenUrl string) bool {
	_, err := url.Parse(givenUrl)
	return err == nil
}

func validDate(givenDate string, mode DateMode) bool {
	var err error
	switch mode {
	case dash:
		_, err = time.Parse("2000-01-01", givenDate)
		break
	case slash:
		_, err = time.Parse("2000/01/01", givenDate)
	}
	return err == nil
}
