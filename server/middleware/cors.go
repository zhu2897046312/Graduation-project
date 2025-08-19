// middleware/cors.go
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 处理跨域请求，支持预检请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有来源（生产环境建议指定具体域名）
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		// 允许的请求方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的请求头（必须包含Authorization，因为你的AuthMiddleware需要它）
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")
		// 允许前端读取的响应头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin")
		// 允许携带cookie
		c.Header("Access-Control-Allow-Credentials", "true")

		// 处理预检请求（OPTIONS）
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent) // 204表示成功但无内容
			return
		}

		c.Next()
	}
}