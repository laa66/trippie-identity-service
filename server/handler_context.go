package httpserver

import (
	"context"

	"github.com/gin-gonic/gin"
)

type HandlerContext interface {
	QueryParam(name string) string
	Param(name string) string
	Header(name string) string
	Context() context.Context
	BindBody(dest any) error
	JSON(status int, obj any)
	Status(status int)
	Error(err error)
}

var _ HandlerContext = (*handlerContext)(nil)

type handlerContext struct {
	Ctx *gin.Context
}

func (h *handlerContext) QueryParam(name string) string {
	return h.Ctx.Query(name)
}

func (h *handlerContext) Param(name string) string {
	return h.Ctx.Param(name)
}

func (h *handlerContext) Header(name string) string {
	return h.Ctx.GetHeader(name)
}

func (h *handlerContext) BindBody(dest any) error {
	return h.Ctx.ShouldBindJSON(dest)
}

func (h *handlerContext) Context() context.Context {
	return h.Ctx.Request.Context()
}

func (h *handlerContext) JSON(status int, obj any) {
	h.Ctx.JSON(status, obj)
}

func (h *handlerContext) Status(status int) {
	h.Ctx.Status(status)
}

func (h *handlerContext) Error(err error) {
	h.Ctx.Error(err)
}
