package websocket

import (
	"IMChat_App/pkg/common"
	"IMChat_App/pkg/protocol"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		MyServer.Ungister <- c
		c.Conn.Close()
	}()

	for {
		// 设置一个 PONG 消息处理函数
		c.Conn.SetPongHandler(func(appData string) error {
			// 收到 PONG 消息后，继续保持连接
			return nil
		})

		// 读取消息
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			MyServer.Ungister <- c
			c.Conn.Close()
			break
		}

		// 处理消息
		msg := &protocol.Message{}
		err = proto.Unmarshal(message, msg)
		if err != nil {
			continue
		}

		// 处理心跳消息
		if msg.Type == common.HEAT_BEAT {
			c.handleHeartbeat()
		} else {
			// 处理其他消息
			MyServer.Broadcast <- message
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	// 将通道中的内容跟发给客户端
	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			break
		}
	}
}

func (c *Client) handleHeartbeat() {
	// 响应心跳消息
	pongMsg := &protocol.Message{
		Content: common.PONG,
		Type:    common.HEAT_BEAT,
	}
	pongBytes, err := proto.Marshal(pongMsg)
	if err != nil {
		return
	}

	err = c.Conn.WriteMessage(websocket.BinaryMessage, pongBytes)
	if err != nil {
		return
	}
}
