package controllers

import (
	"Golang-Docker/pkg/models"
	"Golang-Docker/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/xlzd/gotp"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)
	u := newUser.NewUser()
	if u == nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		result, _ := json.Marshal(u)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	result, _ := json.Marshal(users)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails, _ := models.GetUserById(id)
	if userDetails.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result, _ := json.Marshal(userDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails, db := models.GetUserById(id)
	if userDetails.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if updateUser.FirstName != "" {
		userDetails.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		userDetails.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		userDetails.Password = updateUser.Password
	}
	if updateUser.Role != "" {
		userDetails.Role = updateUser.Role
	}
	db.Save(&userDetails)
	result, _ := json.Marshal(userDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userTemp, _ := models.GetUserById(id)
	if userTemp.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user := models.DeleteUserById(id)
	result, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser = &models.LoginUser{}
	utils.ParseBody(r, loginUser)
	userDetails, _ := models.GetUserByEmail(loginUser.Email)
	if userDetails.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	currentPassword := []byte(userDetails.Password)
	givenPassword := []byte(loginUser.Password)
	err := bcrypt.CompareHashAndPassword(currentPassword, givenPassword)

	// TOTP for LMT4URYNZKEWZRAA - Default secret for the current dependecy repo
	// N.B. secret length is 16, so use gotp.RandomSecret(16)
	totp := gotp.NewDefaultTOTP(userDetails.Secret)
	currentTOTP := totp.At(int(time.Now().Unix()))
	prev30sec := time.Now().Unix() - 30
	next30sec := time.Now().Unix() + 30
	prevTOTP := totp.At(int(prev30sec))
	nextTOTP := totp.At(int(next30sec))

	if (loginUser.Code == currentTOTP || loginUser.Code == prevTOTP || loginUser.Code == nextTOTP) && userDetails.Email == loginUser.Email && err == nil {
		expirationTime := time.Now().Add(1 * time.Hour)
		refreshTime := time.Now().Add(30 * time.Minute)
		claims := &utils.IdentityClaims{
			Username:  userDetails.Email,
			Role:      userDetails.Role,
			ExpiresAt: expirationTime.Unix(),
			RefreshAt: refreshTime.Unix(),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
				Id:        userDetails.Email,
				IssuedAt:  time.Now().Unix(),
				Issuer:    utils.GetIssuer(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		tokenString, err := token.SignedString(utils.GetJWTSecret())
		if err != nil {
			// fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response := &utils.TokenResponse{
			Token:     tokenString,
			Expires:   expirationTime.Unix(),
			Refreshes: refreshTime.Unix(),
		}
		reply, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// fmt.Println(string(reply))
		w.WriteHeader(http.StatusOK)
		w.Write(reply)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func RefreshUser(w http.ResponseWriter, r *http.Request) {

	// safe to run - auth middleware has already run this code
	tokenHeader := r.Header.Get("Authorization")
	splitted := strings.Split(tokenHeader, " ")
	tokenPart := splitted[1]
	claims := &utils.IdentityClaims{}
	token, err := jwt.ParseWithClaims(tokenPart, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetJWTSecret()), nil
	})

	if token.Valid && err == nil {
		expirationTime := time.Now().Add(1 * time.Hour)
		refreshTime := time.Now().Add(30 * time.Minute)
		newClaims := &utils.IdentityClaims{
			Username:  claims.Username,
			Role:      claims.Role,
			ExpiresAt: expirationTime.Unix(),
			RefreshAt: refreshTime.Unix(),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
				Id:        claims.Username,
				IssuedAt:  time.Now().Unix(),
				Issuer:    utils.GetIssuer(),
			},
		}

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, newClaims)
		newTokenString, err := newToken.SignedString(utils.GetJWTSecret())
		if err != nil {
			// fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response := &utils.TokenResponse{
			Token:     newTokenString,
			Expires:   expirationTime.Unix(),
			Refreshes: refreshTime.Unix(),
		}
		reply, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(reply)
		return
	} else {
		response := utils.Message(false, "Malformed authentication token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		utils.Respond(w, response)
		return
	}

}
