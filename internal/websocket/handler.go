package websocket

import (
	api "IMChat_App/api/message"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 用于连接ws
func RunSocket(c *gin.Context) {
	var err error
	var id int
	id_str := c.Query("id")
	if id_str == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account parameter is required"})
		return
	}
	id, err = strconv.Atoi(id_str)

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &Client{
		Id:   int32(id),
		Conn: ws,
		Send: make(chan []byte),
	}

	MyServer.Register <- client

	fmt.Println("websocket connect success")

	go client.Read()
	go client.Write()
}

func GetUserStatus(id int32) bool {
	if _, ok := MyServer.Clients[id]; ok {
		fmt.Println(id, " : online")
		return true
	}
	return false
}

func SendMsg(message []byte) {
	fmt.Println("正在发送mq的消息")

	// 发给hub的消息通道，hub集中处理
	MyServer.Broadcast <- message

	// 保存消息到数据库
	err := api.StoreMsg(message)
	if err != nil {
		fmt.Println("store msg failed")
	}

	// 更新历史消息
	err = api.UpdateHistory(message)
	if err != nil {
		fmt.Println("update msg to redis failed")
	}
}
