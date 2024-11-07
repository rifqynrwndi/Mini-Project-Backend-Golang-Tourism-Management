package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTourism struct {
}

type jwtCustomClaims struct {
	Name   string `json:"name"`
	UserID int    `json:"userID"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (jwtTourism JwtTourism) GenerateJWT(userID int, name, role string) (string, error) {
	claims := &jwtCustomClaims{
		Name:   name,
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}
