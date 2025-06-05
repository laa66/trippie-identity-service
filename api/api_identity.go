package api

import (
	"github.com/gin-gonic/gin"
	apperr "github.com/laa66/trippie-identity-service.git/error"
	http_server "github.com/laa66/trippie-identity-service.git/server"
)

type Object struct {
	Name string
}

func RegisterIdentityEndpoints(rg *gin.RouterGroup) {
	rg.POST("/login", http_server.WrapNoBody(LoginHandler))
}

// TODO: move to handler
func LoginHandler(handlerContext http_server.HandlerContext) (responseCode int, data any, err error) {
	return 0, nil, apperr.New("djwakfkaf").WithHttpStatus(403)
}