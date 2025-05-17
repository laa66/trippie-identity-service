package httpserver

import (
	"context"

	"github.com/gin-gonic/gin"
)

type HandlerContext struct {
	Ctx *gin.Context
}

func (h *HandlerContext) QueryParam(name string) string {
	return h.Ctx.Query(name)
}

func (h *HandlerContext) Param(name string) string {
	return h.Ctx.Param(name)
}

func (h *HandlerContext) Header(name string) string {
	return h.Ctx.GetHeader(name)
}

func (h *HandlerContext) BindBody(dest any) error {
	return h.Ctx.ShouldBindJSON(dest)
}

func (h *HandlerContext) Context() context.Context {
	return h.Ctx.Request.Context()
}

func (h *HandlerContext) JSON(status int, obj any) {
	h.Ctx.JSON(status, obj)
}

func (h *HandlerContext) Status(status int) {
	h.Ctx.Status(status)
}

func (h *HandlerContext) Error(err error) {
	h.Ctx.Error(err)
}