package service

import (
	"github.com/laa66/trippie-identity-service.git/ctx"
	apperr "github.com/laa66/trippie-identity-service.git/error"
	"github.com/laa66/trippie-identity-service.git/internal/core/domain/dto"
)

type (
	IdentityService interface {
		GetIdentity(ctx.Ctx) (*dto.Identity, *apperr.AppErr)
		RegisterIdentity(ctx.Ctx) *apperr.AppErr
	}
)
