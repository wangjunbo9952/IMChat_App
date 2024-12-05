package repository

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/internal/model"
	"github.com/google/uuid"
)

func Register(user *model.User) bool {
	db := pool.GetDB()
	var cnt int64
	db.First(user, "account = ?", user.Account).Count(&cnt)
	if cnt == 0 {
		// 账户不存在可以注册,同时生成uuid
		user.Uuid = uuid.New().String()
		db.Create(user)
		return true
	} else {
		// 账户存在，返回false
		return false
	}

}
