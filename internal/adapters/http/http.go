package http

import (
	"github.com/gin-gonic/gin"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/handlers"
	httpserver "github.com/laa66/trippie-identity-service.git/server"
)

type HTTP struct {
	IdentityHandler handlers.IdentityHandler
}

func NewHTTPServer(ih handlers.IdentityHandler) *HTTP {
	return &HTTP{
		IdentityHandler: ih,
	}
}

func (h *HTTP) RegisterIdentityEndpoints(rg *gin.RouterGroup) {
	rg.GET("/identity", httpserver.WrapNoBody(h.IdentityHandler.GetIdentity))
}
