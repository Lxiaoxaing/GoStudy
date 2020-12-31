package middleware

import "github.com/gin-gonic/gin"

// 设置中间件
func SetMiddleware(key string, value interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, value)
		c.Next()
	}
}
