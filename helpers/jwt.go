package helpers

import (
	"os"
	"time"
	"errors"

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

// Claims represents the JWT claims
type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

// VerifyToken verifies a JWT token and returns the user ID
func VerifyToken(tokenString string, secretKey string) (string, error) {
	// Parse the JWT token string to obtain the token object
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	// Extract the claims and the user ID from the token
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userID := claims.ID
	if userID == "" {
		return "", errors.New("missing user ID in token claims")
	}

	// Check the token expiration time
	if claims.ExpiresAt < time.Now().Unix() {
		return "", errors.New("expired token")
	}

	return userID, nil
}
