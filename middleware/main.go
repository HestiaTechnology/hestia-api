package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqId := uuid.New().String()
		ctx.Writer.Header().Set("X-Request-Id", reqId)
		ctx.Next()
	}
}

func DontCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Cache-Control", "no-store")
		ctx.Next()
	}
}

func MethodNotAllowed() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func NotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}
