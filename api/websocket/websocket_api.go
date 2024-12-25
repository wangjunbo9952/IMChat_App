package websocket

import (
	"IMChat_App/api/message"
	"IMChat_App/internal/websocket"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 连接ws
func ConnectWs(c *gin.Context) {
	websocket.RunSocket(c)
}

// 获取用户在线状态
func GetUserStatus(c *gin.Context) {
	var id int
	id_str := c.Query("id")
	if id_str == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account parameter is required"})
		return
	}
	id, _ = strconv.Atoi(id_str)
	res := websocket.GetUserStatus(int32(id))
	if res == true {
		c.JSON(http.StatusOK, gin.H{"status": "online"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "offline"})
	}
}

// 发送消息
func SendMsg(message []byte) {
	websocket.SendMsg(message)
}

// 获取与该用户的历史消息100条
func GetHistory(c *gin.Context) {
	// Redis实现
	var key string
	message.GetHistoryMsg(key)
}
