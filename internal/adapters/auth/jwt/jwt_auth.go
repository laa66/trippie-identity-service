package jwtauth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID         string
	ExpirationTime time.Time
	IssuedAt       time.Time
	Type           TokenType
	Issuer         string
}

type MapClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

type Tokens map[TokenType]string

type TokenType string

const (
	AccessToken  TokenType = "ACCESS_TOKEN"
	RefreshToken TokenType = "REFRESH_TOKEN"
)

// TODO: specific error if env variable not found, move to util
func GetJWTSecretKey() string {
	return os.Getenv("TRIPPIE_JWT_SECRET_KEY")
}