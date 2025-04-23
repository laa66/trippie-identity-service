package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerWithBody[T any] func(*gin.Context, T) (int, any, error)
type HandlerNoBody func(*gin.Context) (int, any, error)

// TODO: Better error handling, add err wrapper itd and return status code based on err definition
// TODO: Move API to own lib and also errors to own lib
// TODO: Add logger

func WrapWithBody[T any](handler HandlerWithBody[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		if err := c.ShouldBindJSON(&body); err != nil {
			response(c, http.StatusBadRequest, nil, err)
			return
		}

		code, res, err := handler(c, body)

		if res == nil {
			responseNoData(c, code, err)
			return
		}

		response(c, code, res, err)
	}
}

func WrapNoBody(handler HandlerNoBody) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, res, err := handler(c)
		if res == nil {
			responseNoData(c, code, err)
		} else {
			response(c, code, res, err)
		}
	}
}

func responseNoData(c *gin.Context, code int, err error) {
	if err != nil {
		c.JSON(code, err.Error())
	} else {
		c.Status(code)
	}
}

func response(c *gin.Context, code int, data any, err error) {
	if err != nil {
		c.JSON(code, err.Error())
	} else {
		c.JSON(code, data)
	}
}
