package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"simple-REST/pkg/utils"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// List of endpoints that doesn't require auth
		// useless - it's required with the old method
		/*
			notAuth := []string{
				"/public/user/login/",
				"/public/article/page/{pageNumber}",
				"/public/article/{articleId}",
			}
		*/
		requestPath := r.URL.Path // current request path

		// old method - invalid with paths with variables in it (ex. {pageNumber})
		/*
			for _, value := range notAuth {
				if value == requestPath {
					next.ServeHTTP(w, r)
					return
				}
			}
		*/

		// new method
		if strings.Contains(requestPath, "/public/") {
			next.ServeHTTP(w, r)
			return
		}

		// Grab the token from the header and check if Token is missing
		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			response = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// Format `Bearer {token-body}`
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// Grab the token part, what we are truly interested in
		tokenPart := splitted[1]
		tk := &utils.IdentityClaims{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.GetJWTSecret()), nil
		})

		// Malformed token or Token is invalid, maybe not signed on this server
		if err != nil || !token.Valid {
			// fmt.Println(err.Error())
			response = utils.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// Everything went well, proceed with the request and set the caller to the user retrieved
		// from the parsed token and proceed in the middleware chain
		fmt.Sprintf("User %", tk.Username) // Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.Username)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
