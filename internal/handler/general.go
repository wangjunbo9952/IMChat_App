package handler

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadChatPage(c *gin.Context) {
	db := pool.GetDB()

	account := c.Param("account")
	var user model.User
	if err := db.First(&user, "account = ?", account).Error; err != nil {
		return
	}

	// 渲染聊天页面，传入用户数据
	c.HTML(http.StatusOK, "chat.html", gin.H{
		"user": user,
	})

}

func LoadLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"success": true,
	})
}
