package middleware

import (
	"IMChat_App/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AutoRefreshMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查是否需要刷新
		if exp, ok := c.Get("exp"); !ok {
			return
		} else if v, ok := exp.(int64); !ok {
			return
		} else {
			if shouldRefreshToken(v) {
				userID, _ := c.Get("userID")
				userAccount, _ := c.Get("userAccount")
				uid := userID.(uint)
				uac := userAccount.(string)
				newToken, err := service.GenerateJWT(uid, uac)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "刷新token失败"})
					return
				}
				c.Header("New-Token", newToken)
			}
		}

		c.Next()
	}
}

func shouldRefreshToken(exp int64) bool {
	now := time.Now().Unix()

	// 如果剩余时间小于6小时，则刷新
	remainingTime := exp - now
	return remainingTime < 6*3600 // 6小时
}
