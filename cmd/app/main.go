package main

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/internal/kafka"
	"IMChat_App/internal/router"
	"IMChat_App/internal/websocket"
)

func main() {
	err := pool.InitDB()
	if err != nil {
		return
	}

	err = pool.InitRedisClient()
	if err != nil {
		return
	}

	kafka.InitConsumer("chat")
	kafka.InitProducer()
	go kafka.ConsumeMessages(kafka.GetConsumer(), "p2p") // 单聊消息

	go websocket.MyServer.Start()

	router.NewRouter()

}
