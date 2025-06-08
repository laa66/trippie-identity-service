package api

import (
	"github.com/gin-gonic/gin"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/handlers"
	httpserver "github.com/laa66/trippie-identity-service.git/server"
)

type Api struct {
	IdentityHandler handlers.IdentityHandler
}

// TODO: support multiple handlers or factory
func NewApi(ih handlers.IdentityHandler) *Api {
	return &Api{
		IdentityHandler: ih,
	}
}

func (a *Api) RegisterIdentityEndpoints(rg *gin.RouterGroup) {
	rg.GET("", httpserver.WrapNoBody(a.IdentityHandler.GetIdentity))
	rg.POST("", httpserver.WrapWithBody(a.IdentityHandler.RegisterIdentity))
}
