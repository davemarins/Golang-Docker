package models

import (
	"Golang-Docker/pkg/config"
	"github.com/jinzhu/gorm"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var dbUserClient *gorm.DB

type User struct {
	gorm.Model
	// Id string `json:"id"`
	FirstName string `gorm:""json:"first_name"`
	LastName  string `gorm:""json:"last_name"`
	Email     string `gorm:""json:"email"`
	Password  string `gorm:""json:"password"`
	Secret    string `gorm:""json:"secret_totp"`
	Role      string `gorm:""json:"role"`
}

func init() {
	config.Connect()
	dbUserClient = config.GetDB()
	dbUserClient.AutoMigrate(&User{})
}

func (u *User) NewUser() *User {
	policy := bluemonday.UGCPolicy()
	if !validEmail(u.Email) {
		return nil
	}
	u.FirstName = policy.Sanitize(u.FirstName)
	u.LastName = policy.Sanitize(u.LastName)
	u.Secret = policy.Sanitize(u.Secret)
	tempRole := Role(u.Role)
	if tempRole == "" {
		return nil
	}
	u.Role = string(tempRole)
	pwdByte := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hash)
	dbUserClient.NewRecord(u)
	dbUserClient.Create(&u)
	return u
}

func GetAllUsers() []User {
	var users []User
	dbUserClient.Find(&users)
	return users
}

func GetUserById(id int64) (*User, *gorm.DB) {
	var user User
	db := dbUserClient.Where("id = ?", id).Find(&user)
	return &user, db
}

func GetUserByEmail(email string) (*User, *gorm.DB) {
	var user User
	db := dbUserClient.Where("email = ?", email).Find(&user)
	return &user, db
}

func DeleteUserById(id int64) User {
	var user User
	dbUserClient.Where("id = ?", id).Delete(user)
	return user
}
