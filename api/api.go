package api

import "github.com/gin-gonic/gin"

type Api interface {
	SetupRouter() *gin.Engine
}

type api struct {}

func NewApi() *api {
	return &api{}
}