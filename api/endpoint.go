package api

import (
	"github.com/gin-gonic/gin"
	appCtx "github.com/laa66/trippie-identity-service.git/ctx"
)

// Interfejs ogólny
type Endpoint interface {
	Method() string
	Path() string
	Handler() gin.HandlerFunc
}

// Generyczna struktura endpointu
type GenericEndpoint[T any] struct {
	MethodStr string
	PathStr   string
	HandlerFn func(ctx appCtx.AppContext, body T) (int, any, error)
}

func (e GenericEndpoint[T]) Method() string {
	return e.MethodStr
}

func (e GenericEndpoint[T]) Path() string {
	return e.PathStr
}

func (e GenericEndpoint[T]) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		code, res, err := e.HandlerFn(appCtx.AppContext{Context: c}, body)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(code, res)
	}
}
