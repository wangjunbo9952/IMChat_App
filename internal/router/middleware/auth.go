package middleware

import (
	"IMChat_App/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 定义常量
const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
)

// AuthMiddleware 用于验证JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		// 检查JWT格式，Authorization: Bearer <token>
		if !strings.HasPrefix(authHeader, BearerPrefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// 提取Token
		tokenString := authHeader[len(BearerPrefix):]

		// 验证JWT
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 将用户ID存储到请求的Context中，供后续的处理使用
		userID := claims.Subject
		userAccount := claims.Audience
		exp := claims.ExpiresAt

		// 将用户信息存储到Gin Context中
		c.Set("userID", userID)
		c.Set("userAccount", userAccount)
		c.Set("exp", exp)

		// 继续执行后续的请求处理
		c.Next()
	}
}
