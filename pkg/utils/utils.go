package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// IMPORTANT: You need to export the fields in all objects by capitalizing the first letter in the field name.
// Example: in TokenResponse no "token string" but "Token string"

type IdentityClaims struct {
	Username  string `json:"username"`
	Role      string `json:"role"`
	ExpiresAt int64  `json:"expires_at"`
	RefreshAt int64  `json:"refresh_at"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Token     string
	Expires   int64
	Refreshes int64
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func GetJWTSecret() []byte {
	// put your secret here
	return []byte("oi2K2wax3RX7Gi6bef7PBR2OuZQt1UMUDTnKiqC5JCsToMbfBUgh52l2e9CRyF7AjZnCXMkG8o6mR4E3n2QyktBUZ7unyBVx0Ai")
}

func GetIssuer() string {
	// put your issuer name for jwt here
	return "simple-REST"
}
