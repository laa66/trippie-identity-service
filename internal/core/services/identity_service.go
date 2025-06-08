package services

import (
	"time"

	"github.com/laa66/trippie-identity-service.git/ctx"
	apperr "github.com/laa66/trippie-identity-service.git/error"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/repository"
	"github.com/laa66/trippie-identity-service.git/internal/core/domain/dto"
	"github.com/laa66/trippie-identity-service.git/internal/core/domain/entity"
	"github.com/laa66/trippie-identity-service.git/internal/core/ports/service"
)

var _ service.IdentityService = (*identityService)(nil)

type identityService struct {
	Repositories repository.PostgresRepositories
}

func NewIdentityService(repositories repository.PostgresRepositories) *identityService {
	return &identityService{
		Repositories: repositories,
	}
}

func (i *identityService) GetIdentity(ctx ctx.Ctx) (*dto.Identity, *apperr.AppErr) {
	logger.Log().Debug("enter get identity")
	return nil, nil
}

func (i *identityService) RegisterIdentity(ctx ctx.Ctx, identity *dto.CreateIdentity) *apperr.AppErr {
	logger.Log().Debug("register identity")
	err := i.Repositories.GetIdentityRepository().Create(&entity.Identity{
		Mail: identity.Mail,
		Password: identity.Password,
		Date: time.Now(),
	})

	return err
}
