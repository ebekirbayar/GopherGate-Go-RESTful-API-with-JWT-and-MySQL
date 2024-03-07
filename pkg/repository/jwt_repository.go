package repository

import (
	"GopherGate/pkg/middleware"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// CreateToken generates a JWT token with the provided expiration duration, user name, user ID, and user role.
func CreateToken(userName string, userId uint, userRole string) (*string, error) {
	privateKey := []byte("GopherGate")

	// Token süresi 1 gün olarak ayarlanır
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create a token with the private key
	claims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: expirationTime.Unix(),
	}

	// Convert uint to string
	userIdStr := strconv.Itoa(int(userId))

	// Create JWT with custom claims
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JwtCustom{
		StandardClaims: claims,
		UserID:         userIdStr,
		UserName:       userName,
		UserRole:       userRole,
	})

	// Sign the token
	tokenString, err := accessToken.SignedString(privateKey)
	if err != nil {
		return nil, errors.New(err.Error() + " error occurred while signing JWT!")
	}
	return &tokenString, nil
}
