package handler

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadChatPage(c *gin.Context) {
	db := pool.GetDB()

	account := c.Param("account")
	var user model.User
	if err := db.First(&user, "account = ?", account).Error; err != nil {
		fmt.Println(account)
		return
	}

	// 渲染聊天页面，传入用户数据
	c.HTML(http.StatusOK, "chat.html", gin.H{
		"user": user,
	})

}
