package pool

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	Ctx = context.Background()
	Rds *redis.Client
)

func InitRedisClient() error {
	Rds := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})

	// 测试连接
	pong, err := Rds.Ping(Ctx).Result()
	if err != nil {
		fmt.Printf("连接Redis出错，错误信息：%v", err)
		return err
	}
	fmt.Println("成功连接redis", pong)
	return nil
}
