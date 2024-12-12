package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocket(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User parameter is required"})
		return
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &Client{
		Name: user,
		Conn: ws,
		Send: make(chan []byte),
	}

	MyServer.Register <- client

	fmt.Println("websocket connect success")

	go client.Read()
	go client.Write()
}
