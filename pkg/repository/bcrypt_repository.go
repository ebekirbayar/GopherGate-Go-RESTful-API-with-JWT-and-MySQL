package repository

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given user password using bcrypt.
func HashPassword(userPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash checks if the given user password matches the hashed password.
func CheckPasswordHash(userPassword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(userPassword))
	return err == nil
}
