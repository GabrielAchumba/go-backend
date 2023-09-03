package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type appHandler func(ctx *gin.Context) *Response

func ServeHTTP(handle appHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := handle(ctx)
		if result == nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success:    false,
				Message:    "INTERNAL SERVER ERROR",
				Data:       nil,
				StatusCode: http.StatusInternalServerError,
			})
		} else {
			ctx.JSON(result.StatusCode, result)
		}
	}
}
