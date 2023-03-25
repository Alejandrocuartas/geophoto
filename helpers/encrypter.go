package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// create a function that encrypts a password and fill it
func EncryptPassword(password string) (string, error) {

	//create a byte array
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// create a function that validate the password and return true if it succeeded
func ValidatePassword(password string) (bool, error) {

	//create a byte array
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	isCorrect := bcrypt.CompareHashAndPassword(bytes, []byte(password)) == nil
	return isCorrect, nil
}
