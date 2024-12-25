package main

import (
	"IMChat_App/internal/storage"
	"IMChat_App/pkg/redis"
	"fmt"
)

func main() {
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
	fmt.Println("初始化成功")
}
