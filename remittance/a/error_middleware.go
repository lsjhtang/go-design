package a

import "github.com/gin-gonic/gin"

func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := recover(); err != nil {
			context.AbortWithStatusJSON(400, gin.H{"error": err})
		} else {
			context.Next() //继续往下走
		}
	}
}
