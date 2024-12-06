package repository

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/internal/model"
	"github.com/google/uuid"
)

func Register(user *model.User) bool {
	db := pool.GetDB()
	var cnt int64
	result := db.Model(user).Where("account = ?", user.Account).Count(&cnt)
	if result.Error != nil {
		return false
	}
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

func Login(user *model.User) bool {
	db := pool.GetDB()

	var recUser model.User
	db.First(&recUser, "account = ?", user.Account)
	if recUser.Password != user.Password {
		// 账户不存在
		return false
	}
	return true
}

func SearchUser(account string) (bool, *model.User) {
	db := pool.GetDB()

	var user model.User
	db.First(&user, "account = ?", account)

	if user.Account != account {
		return false, nil
	}

	return true, &user

}
