package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	apperr "github.com/laa66/trippie-identity-service.git/error"
	jwtauth "github.com/laa66/trippie-identity-service.git/internal/adapters/auth/jwt"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
)

var _ JWTService = (*jwtService)(nil)

type JWTService interface {
	GenerateTokens(userID string) (jwtauth.Tokens, *apperr.AppErr)
	VerifyToken(token string) (*jwtauth.Claims, *apperr.AppErr)
	RefreshToken(token string) (jwtauth.Tokens, *apperr.AppErr)
}

type jwtService struct {
	secret string
}

func NewJWTService(secret string) *jwtService {
	return &jwtService{
		secret: secret,
	}
}

func (j *jwtService) GenerateTokens(userID string) (jwtauth.Tokens, *apperr.AppErr) {
	logger.Log().Debug("enter generate tokens")

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(15 * time.Minute).UnixNano(),
		"iat": time.Now().UnixNano(),
		"iss": "trippie",
	}

	logger.Log().Debug("created token claims", "claims", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	logger.Log().Debug("new token", "JWT", token)

	accessToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		logger.Log().Error("error while signing JWT token", "error", err)
		return nil, apperr.Wrap(err)
	}
	logger.Log().Debug("signed access token", "JWT", accessToken)

	return map[jwtauth.TokenType]string{
		jwtauth.AccessToken: accessToken,
		// RefreshToken: refreshToken,
	}, nil
}

func (j *jwtService) VerifyToken(tokenStr string) (*jwtauth.Claims, *apperr.AppErr) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil || !token.Valid {
		return nil, apperr.Wrap(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, apperr.Wrap(err)
	}

	return &jwtauth.Claims{
		UserID:         claims["sub"].(string),
		ExpirationTime: time.Unix(int64(claims["exp"].(float64)),0),
		IssuedAt:       time.Unix(int64(claims["iat"].(float64)),0),
		Issuer:         claims["iss"].(string),
	}, nil
}

func (j *jwtService) RefreshToken(tokenStr string) (jwtauth.Tokens, *apperr.AppErr) {
	panic("unimplemented")
}
