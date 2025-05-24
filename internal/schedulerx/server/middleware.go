package server

import (
	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, x-wellness-jwt, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(200)
			return
		}

		context.Next()
	}
}
