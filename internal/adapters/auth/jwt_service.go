package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	apperr "github.com/laa66/trippie-identity-service.git/error"
	jwtauth "github.com/laa66/trippie-identity-service.git/internal/adapters/auth/jwt"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
)

var _ JWTService = (*jwtService)(nil)

type JWTService interface {
	GenerateTokens(userID string, tokens ...jwtauth.TokenType) (jwtauth.Tokens, *apperr.AppErr)
	VerifyToken(token string) (*jwtauth.Claims, *apperr.AppErr)
	RefreshToken(token string) (jwtauth.Tokens, *apperr.AppErr)
}

type jwtService struct{}

func NewJWTService() *jwtService {
	return &jwtService{}
}

func (j *jwtService) GenerateTokens(userID string, tokenTypes ...jwtauth.TokenType) (jwtauth.Tokens, *apperr.AppErr) {
	logger.Log().Debug("enter generate tokens", "userID", userID, "token types", tokenTypes)

	tokens := jwtauth.Tokens{}
	for _, tokenType := range tokenTypes {
		token, err := j.CreateToken(userID, tokenType, jwtauth.GetJWTSecretKey())
		if err != nil {
			logger.Log().Error("error while creating JWT token", "error", err)
			return nil, err
		}
		tokens[tokenType] = *token
	}

	return tokens, nil
}

func (j *jwtService) CreateToken(userID string, tokenType jwtauth.TokenType, secretKey string) (*string, *apperr.AppErr) {
	claims := jwtauth.MapClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       uuid.NewString(),
			Subject:  userID,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer:   "trippie",
		},
		Role: string(tokenType),
	}

	//TODO: export to config
	switch tokenType {
	case jwtauth.AccessToken:
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(15 * time.Minute))
	case jwtauth.RefreshToken:
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(60 * time.Minute))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := jwtauth.GetJWTSecretKey()
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.Log().Error("error while signing JWT token", "error", err)
		return nil, apperr.Wrap(err)
	}
	return &signedToken, nil
}

func (j *jwtService) VerifyToken(tokenStr string) (*jwtauth.Claims, *apperr.AppErr) {
	logger.Log().Debug("enter verify token", "token", tokenStr)
	secret := jwtauth.GetJWTSecretKey()
	claims := &jwtauth.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, apperr.Wrap(err)
	}
	logger.Log().Debug("parsed token with claims", "token", token, "claims", claims)

	return &jwtauth.Claims{
		UserID:         claims.Subject,
		ExpirationTime: claims.ExpiresAt.Time,
		IssuedAt:       claims.IssuedAt.Time,
		Issuer:         claims.Issuer,
		Type:           jwtauth.TokenType(claims.Role),
	}, nil
}

func (j *jwtService) RefreshToken(refreshToken string) (jwtauth.Tokens, *apperr.AppErr) {
	logger.Log().Debug("enter refresh token", "token", refreshToken)
	claims, err := j.VerifyToken(refreshToken)
	if err != nil {
		logger.Log().Error("error while veryfing token", "error", err)
		return nil, err
	}

	tokens, err := j.GenerateTokens(claims.UserID, jwtauth.AccessToken, jwtauth.RefreshToken)
	if err != nil {
		logger.LogErr(err)
		return nil, err
	}

	return tokens, nil
}
