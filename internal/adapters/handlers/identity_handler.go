package handlers

import (
	nethttp "net/http"

	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"github.com/laa66/trippie-identity-service.git/internal/core/domain/dto"
	"github.com/laa66/trippie-identity-service.git/internal/core/ports/service"
	httpserver "github.com/laa66/trippie-identity-service.git/server"
)

type IdentityHandler interface {
	GetIdentity(handlerContext httpserver.HandlerContext) (responseCode int, data any, err error)
	RegisterIdentity(handlerContext httpserver.HandlerContext, identity *dto.CreateIdentity) (responseCode int, data any, err error)
}

var _ IdentityHandler = (*identityHandler)(nil)

type identityHandler struct {
	IdentityService service.IdentityService
}

func NewIdentityHandler(identityService service.IdentityService) *identityHandler {
	return &identityHandler{
		IdentityService: identityService,
	}
}

func (i *identityHandler) GetIdentity(handlerContext httpserver.HandlerContext) (responseCode int, data any, err error) {
	// TODO: get identity based on identifiers in ctx
	identity, err := i.IdentityService.GetIdentity(handlerContext)
	return nethttp.StatusOK, identity, err
}

func (i *identityHandler) RegisterIdentity(handlerContext httpserver.HandlerContext, identity *dto.CreateIdentity) (responseCode int, data any, err error) {
	apperr := i.IdentityService.RegisterIdentity(handlerContext, identity)
	if apperr != nil {
		logger.Log().Debug("register identity", "error", apperr)
		return 0, nil, apperr
	}

	logger.Log().Debug("register identity success")
	return nethttp.StatusOK, nil, nil
}
