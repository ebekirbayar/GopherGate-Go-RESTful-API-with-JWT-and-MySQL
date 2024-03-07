package middleware

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// JwtCustom defines the custom claims for JWT tokens.
type JwtCustom struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	UserRole string `json:"userRole"`
	jwt.StandardClaims
}

// Auth is a middleware function that verifies JWT token and checks its content.
func Auth(ctx *fiber.Ctx) error {
	// Get and verify JWT token
	accessToken, err := getTokenFromHeader(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Unauthorized", "error": err.Error()})
	}

	// If verification is successful, set user identity in Fiber context
	ctx.Locals("userID", accessToken.UserID)
	ctx.Locals("userName", accessToken.UserName)
	ctx.Locals("userRole", accessToken.UserRole)

	// Continue to the next middleware or handler
	return ctx.Next()
}

// getTokenFromHeader retrieves and verifies JWT token from Authorization header.
func getTokenFromHeader(c *fiber.Ctx) (*JwtCustom, error) {
	// Retrieve token from Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		fmt.Println("Authorization header not found")
		return nil, errors.New("authorization header not found")
	}

	// Extract and verify token
	tokenString := getTokenStringFromHeader(authHeader)
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustom{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Invalid signing method")
			return nil, errors.New("invalid signing method")
		}
		return []byte("GopherGate"), nil
	})

	// Handle errors
	if err != nil {
		fmt.Println("ParseWithClaims error:", err)
		return nil, err
	}
	if !token.Valid {
		fmt.Println("Invalid accessToken")
		return nil, errors.New("invalid accessToken")
	}

	// Verify token content
	claims, ok := token.Claims.(*JwtCustom)
	if !ok {
		fmt.Println("Invalid claims")
		return nil, errors.New("invalid claims")
	}

	// If content verification is successful, return the JWT token
	return claims, nil
}

// getTokenStringFromHeader extracts JWT token from Authorization header.
func getTokenStringFromHeader(authHeader string) string {
	// Authorization header should be in the format: Bearer <token>
	const prefix = "Bearer "
	if len(authHeader) > len(prefix) && authHeader[:len(prefix)] == prefix {
		return authHeader[len(prefix):]
	}
	return ""
}
