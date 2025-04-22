package api

import (
	"github.com/gin-gonic/gin"
)

type Api struct {
	Router *gin.Engine
}

func NewApi() *Api {
	return &Api{Router: gin.Default()}
}

func (a *Api) RegisterEndpoints(endpoints []Endpoint) {
	for _, e := range endpoints {
		a.Router.Handle(e.Method(), e.Path(), e.Handler())
	}
}

func (a *Api) Run() {
	a.Router.Run(":8080")
}
