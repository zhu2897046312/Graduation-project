package middleware

import (
	"net/http"
	"strings"

	"server/utils" // 替换为你的实际项目路径

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// AuthMiddleware 认证中间件
func AuthMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头或cookie中获取token
		token := extractToken(c)

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "未提供认证令牌",
			})
			return
		}

		// 从Redis中获取用户信息
		user, err := utils.GetSession(rdb, token, nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "认证服务异常",
			})
			return
		}

		if user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "认证令牌无效或已过期",
			})
			return
		}

		// 将用户信息存入上下文
		c.Set("user", user)
		c.Set("userID", user.ID)
		c.Set("token", token)

		c.Next()
	}
}

// OptionalAuthMiddleware 可选认证中间件（不强制要求登录）
func OptionalAuthMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)

		if token != "" {
			user, err := utils.GetSession(rdb, token, nil)
			if err == nil && user != nil {
				c.Set("user", user)
				c.Set("userID", user.ID)
				c.Set("token", token)
			}
		}

		c.Next()
	}
}

// AuthMiddleware 认证中间件（JWT版本）
func AuthClientMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中提取token
		token := extractToken(c)

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "未提供认证令牌",
			})
			return
		}

		// 解析JWT token
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "认证令牌无效或已过期",
			})
			return
		}

		// 将用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("token", token)

		c.Next()
	}
}

// OptionalAuthMiddleware 可选认证中间件（JWT版本）
func OptionalClientAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)

		if token != "" {
			claims, err := utils.ParseToken(token)
			if err == nil {
				c.Set("userID", claims.UserID)
				c.Set("email", claims.Email)
				c.Set("token", token)
			}
		}

		c.Next()
	}
}

// extractToken 从请求中提取token
// extractToken 从请求中提取token
// extractToken 从请求中提取token
func extractToken(c *gin.Context) string {
	// 1. 从Authorization头获取（支持带Bearer前缀和不带前缀两种情况）
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		// 检查是否包含Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
		// 如果没有Bearer前缀，直接使用整个Authorization头作为token
		return authHeader
	}

	// 2. 从查询参数获取
	if token := c.Query("token"); token != "" {
		return token
	}

	// 3. 从特定Cookie获取 (YEX_AUTH)
	if cookie, err := c.Cookie("YEX_AUTH"); err == nil {
		return cookie
	}

	// 4. 从通用Cookie获取 (兼容旧代码)
	if cookie, err := c.Cookie("auth_token"); err == nil {
		return cookie
	}

	return ""
}

// GetUserFromContext 从上下文中获取用户信息
func GetUserFromContext(c *gin.Context) *utils.CoreLoginUserInfoModel {
	user, exists := c.Get("user")
	if !exists {
		return nil
	}

	if userInfo, ok := user.(*utils.CoreLoginUserInfoModel); ok {
		return userInfo
	}

	return nil
}

// GetUserIDFromContext 从上下文中获取用户ID
func GetUserIDFromContext(c *gin.Context) int64 {
	userID, exists := c.Get("userID")
	if !exists {
		return 0
	}

	if id, ok := userID.(int64); ok {
		return id
	}

	return 0
}

// GetTokenFromContext 从上下文中获取token
func GetTokenFromContext(c *gin.Context) string {
	token, exists := c.Get("token")
	if !exists {
		return ""
	}

	if tokenStr, ok := token.(string); ok {
		return tokenStr
	}

	return ""
}


// DeviceFingerprintMiddleware 设备指纹中间件
func DeviceFingerprintMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取设备指纹
		deviceFingerprint := c.GetHeader("X-Device-Fingerprint")
		
		// 如果头信息中没有，尝试从查询参数获取
		if deviceFingerprint == "" {
			deviceFingerprint = c.Query("device_fingerprint")
		}

		// 如果还是没有，尝试从cookie获取
		if deviceFingerprint == "" {
			if cookie, err := c.Cookie("device_fingerprint"); err == nil {
				deviceFingerprint = cookie
			}
		}

		// 将设备指纹存入上下文（无论是否为空）
		c.Set("device_fingerprint", deviceFingerprint)
		
		// 设置响应头，方便前端使用
		if deviceFingerprint != "" {
			c.Header("X-Device-Fingerprint", deviceFingerprint)
		}

		c.Next()
	}
}

// GetDeviceFingerprintFromContext 从上下文中获取设备指纹
func GetDeviceFingerprintFromContext(c *gin.Context) string {
	fingerprint, exists := c.Get("device_fingerprint")
	if !exists {
		return ""
	}

	if fp, ok := fingerprint.(string); ok {
		return fp
	}

	return ""
}