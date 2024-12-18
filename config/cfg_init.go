package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var JwtSecret string

func init() {
	// 设置文件名
	viper.SetConfigName("config")
	// 设置文件类型
	viper.SetConfigType("yaml")
	// 设置文件路径，可以多个viper会根据设置顺序依次查找
	viper.SetConfigFile("config\\config.yaml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	JwtSecret = viper.GetString("jwt.secretKey")
}
