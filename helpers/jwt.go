package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken generates a new JWT token for the given user ID
func GenerateToken(userID string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// we want to include in the token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
