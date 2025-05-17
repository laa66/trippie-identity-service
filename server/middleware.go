package httpserver

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	apperr "github.com/laa66/trippie-identity-service.git/error"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("panic recovered", "error", r)
				c.AbortWithStatusJSON(http.StatusInternalServerError, apperr.New("internal server error").WithHttpStatus(500))
			}
		}()

		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		var appErr *apperr.AppErr
		if errors.As(err.Err, &appErr) {
			fmt.Printf("%+v\n", appErr.WrappedError())
			c.JSON(appErr.Code, appErr)
			return
		}

		wrapped := apperr.Wrap(err.Err)
		fmt.Printf("%+v\n", wrapped.WrappedError())
		c.JSON(500, wrapped)
	}
}