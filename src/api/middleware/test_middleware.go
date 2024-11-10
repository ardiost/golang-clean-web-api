package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apikey := ctx.GetHeader("x-api-key")
		if apikey == "1" {
			ctx.Next()
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": "apikey is required",
		})
		return
	}
}
