package services

import (
	"time"

	"github.com/laa66/trippie-identity-service.git/ctx"
	apperr "github.com/laa66/trippie-identity-service.git/error"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/auth"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/repository"
	"github.com/laa66/trippie-identity-service.git/internal/core/domain/dto"
	"github.com/laa66/trippie-identity-service.git/internal/core/domain/entity"
	"github.com/laa66/trippie-identity-service.git/internal/core/ports/service"
)

var _ service.IdentityService = (*identityService)(nil)

type identityService struct {
	AuthService  auth.JWTService
	Repositories repository.PostgresRepositories
}

func NewIdentityService(repositories repository.PostgresRepositories, jwtService auth.JWTService) *identityService {
	return &identityService{
		AuthService: jwtService,
		Repositories: repositories,
	}
}

func (i *identityService) GetIdentity(ctx ctx.Ctx) (*dto.Identity, *apperr.AppErr) {
	logger.Log().Debug("enter get identity")
	claims, err := i.AuthService.VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDk5MTcxMTcsImlhdCI6MTc0OTkxNjIxNywiaXNzIjoidHJpcHBpZSIsInN1YiI6IjEifQ.8Y5F0MMd-YqfgVbMwu_SzzD63blgqvD6FYC1AUCMz2I")
	if err != nil {
		logger.Log().Error("error veryfing token", "error", err)
		return nil, err
	}
	logger.Log().Debug("claims", "claims", claims)
	return nil, nil
}

func (i *identityService) RegisterIdentity(ctx ctx.Ctx, identity *dto.CreateIdentity) *apperr.AppErr {
	logger.Log().Debug("register identity")
	tokens, err := i.AuthService.GenerateTokens("1")
	if err != nil {
		logger.Log().Error("get identity error while generating tokens", "error", err)
		return err
	}

	logger.Log().Debug("generated tokens", "tokens", tokens)

	err = i.Repositories.GetIdentityRepository().Create(&entity.Identity{
		Mail:     identity.Mail,
		Password: identity.Password,
		Date:     time.Now(),
	})
	return err
}
