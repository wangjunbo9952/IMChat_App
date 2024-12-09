package handler

import (
	"IMChat_App/internal/model"
	"IMChat_App/internal/service"
	"IMChat_App/pkg/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	// 密码有待完善，掩盖
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "注册失败",
		})
		return
	}
	res := service.Register(&user)
	if res == false {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "注册失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "登录失败",
		})
		return
	}
	res := service.Login(&user)
	if res == false {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "登录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"account": user.Account,
		"message": "登录成功",
	})
}

func SearchUser(c *gin.Context) {
	var account string
	account = c.Query("account")
	flag, res := service.SearchUser(account)
	if flag == false {
		fmt.Println("查找用户不存在")
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "查找用户不存在",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"id":       res.ID,
		"account":  res.Account,
		"username": res.Username,
		"avatar":   res.Avatar,
	})
}

func AddUser(c *gin.Context) {
	var req common.AddUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的请求数据",
		})
		return
	}

	service.AddUser(&req)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "好友请求已发送",
	})
}

func GetFriendsList(c *gin.Context) {
	userId := c.Query("userId")

	userIdUint, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的用户ID",
		})
		return
	}

	finalUserId := uint(userIdUint)

	var res *[]model.User
	res = service.GetFriendsList(finalUserId)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"friends": *res,
	})
}
