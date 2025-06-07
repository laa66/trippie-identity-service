package services

import (
	"github.com/laa66/trippie-identity-service.git/ctx"
	apperr "github.com/laa66/trippie-identity-service.git/error"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"github.com/laa66/trippie-identity-service.git/internal/core/domain/dto"
	"github.com/laa66/trippie-identity-service.git/internal/core/ports/service"
)

var _ service.IdentityService = (*identityService)(nil)

type identityService struct {
	
}

func NewIdentityService() *identityService {
	return &identityService{}
}


func (i *identityService) GetIdentity(ctx ctx.Ctx) (*dto.Identity, *apperr.AppErr) {
	logger.Log().Debug("enter get identity")
	return nil, nil
}

func (i *identityService) RegisterIdentity(ctx.Ctx) *apperr.AppErr {
	panic("unimplemented")
}