package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func InitRedisClient() error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "122.51.195.185:6379",
		Password: "123123",
		DB:       0,
	})

	// 测试连接
	pong, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Printf("连接Redis出错，错误信息：%v", err)
		return err
	}
	fmt.Println("成功连接redis", pong)
	return nil
}
