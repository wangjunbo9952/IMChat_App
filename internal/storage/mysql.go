package storage

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitDB() error {
	// 配置数据库连接信息
	dsn := "root:123123@tcp(localhost:3306)/imchat?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接 MySQL 数据库
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("MySQL start to work ...")

	return nil
}
