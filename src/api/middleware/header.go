package middleware

import "github.com/gin-gonic/gin"

// Header 头处理
func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		c.Writer.Header().Set("Server", "W4WServer")
	}
}
