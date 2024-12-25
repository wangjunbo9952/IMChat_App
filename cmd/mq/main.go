package main

import (
	"IMChat_App/api/websocket"
	"IMChat_App/internal/mq/rabbit"
	"fmt"
)

func main() {
	// 初始化一个最大为10的连接池
	rabbit.NewConnectionPool(10)
	fmt.Println("pool init success")

	// 协程启动消费，持续监听消息队列
	go rabbit.ConsumeMessages(websocket.SendMsg)
}
