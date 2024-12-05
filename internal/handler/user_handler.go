package handler

import (
	"IMChat_App/internal/model"
	"IMChat_App/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	// 密码有待完善，掩盖
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("bind failed!")
		c.JSON(400, gin.H{"message": "fail"})
		return
	}
	res := service.Register(&user)
	if res == false {
		c.JSON(http.StatusOK, gin.H{"message": "fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
