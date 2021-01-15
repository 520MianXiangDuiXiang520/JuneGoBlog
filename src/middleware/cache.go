package middleware

import (
	"github.com/gin-gonic/gin"
)

func CacheMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Cache-Control", "public,max-age=3600")
	}
}
