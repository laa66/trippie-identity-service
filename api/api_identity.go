package api

import (
	"github.com/gin-gonic/gin"
)

type Object struct {
	Name string
}

func (a *api) RegisterIdentityEndpoints(rg *gin.RouterGroup) *api {
	rg.POST("/login", WrapNoBody(LoginHandler))
	return a
}


// TODO: move to handler
func LoginHandler(ctx *gin.Context) (int, any, error) {
	return 200, Object{
		Name: "mx",
	}, nil
}