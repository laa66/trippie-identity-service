package api

import (
	"github.com/gin-gonic/gin"
	http_server "github.com/laa66/trippie-identity-service.git/server"
)

type Object struct {
	Name string
}

func RegisterIdentityEndpoints(rg *gin.RouterGroup) {
	rg.POST("/login", http_server.WrapNoBody(LoginHandler))
}

// TODO: move to handler
func LoginHandler(ctx *gin.Context) (int, any, error) {
	return 200, Object{
		Name: "mx",
	}, nil
}