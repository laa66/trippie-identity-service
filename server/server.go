package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	apperr "github.com/laa66/trippie-identity-service.git/error"
)

// TODO: move to lib
type HttpServer struct {
	Engine *gin.Engine
}

func NewHttpServer(engine *gin.Engine) *HttpServer {
	return &HttpServer{
		Engine: engine,
	}
}

func (h *HttpServer) GetRouterGroup(prefix string) *gin.RouterGroup {
	return h.Engine.Group(prefix)
}

func (h *HttpServer) Run(port int) {
	// TODO: move to config
	h.Engine.Run(fmt.Sprintf(":%d", port))
}

type HandlerWithBody[T any] func(HandlerContext, T) (int, any, error)
type HandlerNoBody func(HandlerContext) (int, any, error)

func WrapWithBody[T any](handler HandlerWithBody[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		if err := c.ShouldBindJSON(&body); err != nil {
			response(c, http.StatusBadRequest, nil, err)
			return
		}

		handlerContext := &handlerContext{Ctx: c}

		code, res, err := handler(handlerContext, body)

		if res == nil {
			responseNoData(c, code, err)
			return
		}

		response(c, code, res, err)
	}
}

func WrapNoBody(handler HandlerNoBody) gin.HandlerFunc {
	return func(c *gin.Context) {
		handlerContext := &handlerContext{Ctx: c}
		code, res, err := handler(handlerContext)
		fmt.Printf("handler response: {code: %d, res: %s, err: %+v}\n", code, res, err)
		if res == nil {
			responseNoData(c, code, err)
		} else {
			response(c, code, res, err)
		}
	}
}

func responseNoData(c *gin.Context, code int, err error) {
	if err != nil {
		c.Error(err)
	} else {
		c.Status(code)
	}
}

func response(c *gin.Context, code int, data any, err error) {
	if err != nil {
		if apperr, ok := err.(*apperr.AppErr); ok {
			if apperr == nil {
				c.JSON(code, data)
				return
			}
			c.Error(apperr)
			return
		}
		c.Error(err)
		return
	}
	c.JSON(code, data)
}
