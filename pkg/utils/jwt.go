package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_secret_key")

func GenerateToken(userID string) (string, error) {

	// Get the token expiration hours from environment variable
	expirationHours, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_HOURS"))
	if err != nil {
		// If conversion fails, set a default expiration time
		expirationHours = 72 // default to 72 hours
	}

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(expirationHours)).Unix(),
		Issuer:    "your_app_name",
		Subject:   userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
