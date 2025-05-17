package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *HttpServer) Run() {
	// TODO: move to config
	h.Engine.Run(":8080")
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

		handlerContext := HandlerContext{Ctx: c}

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
		handlerContext := HandlerContext{Ctx: c}
		code, res, err := handler(handlerContext)
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
		c.Error(err)
	} else {
		c.JSON(code, data)
	}
}
