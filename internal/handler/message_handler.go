package handler

import (
	"IMChat_App/internal/service"
	"IMChat_App/pkg/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMsg(c *gin.Context) {
	var msgReq common.MsgReq
	err := c.ShouldBindJSON(&msgReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的MsgReq",
		})
		return
	}
	fmt.Printf("%v", msgReq)

	reps := service.GetMsg(&msgReq)
	fmt.Printf("%v", reps)
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"messages": reps,
	})
}
