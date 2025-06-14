package jwtauth

import "time"

type Claims struct {
	UserID         string
	ExpirationTime time.Time
	IssuedAt       time.Time
	Type           TokenType
	Issuer         string
}

type Tokens map[TokenType]string

type TokenType string

const (
	AccessToken  TokenType = "ACCESS_TOKEN"
	RefreshToken TokenType = "REFRESH_TOKEN"
)
