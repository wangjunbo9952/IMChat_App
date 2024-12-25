package main

import (
	api "IMChat_App/api/websocket"
	"IMChat_App/internal/mq/rabbit"
	"IMChat_App/internal/router"
	"IMChat_App/internal/storage"
	inter "IMChat_App/internal/websocket"
	"IMChat_App/pkg/redis"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// api.NewWebSocketRouter(r)
	go inter.MyServer.Start()

	// 以下代码测试用
	// 初始化一个最大为10的连接池
	rabbit.NewConnectionPool(10)
	fmt.Println("pool init success")

	// 协程启动消费，持续监听消息队列
	go rabbit.ConsumeMessages(api.SendMsg)

	// 初始化MySQL
	err := storage.InitDB()
	if err != nil {
		return
	}

	// 初始化Redis
	err = redis.InitRedisClient()
	if err != nil {
		return
	}

	router.NewRouter()

	err = r.Run(":9091")
	if err != nil {
		fmt.Println("Run socket error :", err)
	}

}
