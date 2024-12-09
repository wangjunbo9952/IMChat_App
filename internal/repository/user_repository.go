package repository

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/internal/model"
	"IMChat_App/pkg/common"
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

func AddUser(relation *common.AddUserReq) {
	db := pool.GetDB()
	var fri_ship model.FriendShip
	if relation.UserID <= relation.FriendID {
		fri_ship.UserId = uint(relation.UserID)
		fri_ship.FriendId = uint(relation.FriendID)
	} else {
		fri_ship.UserId = uint(relation.FriendID)
		fri_ship.FriendId = uint(relation.UserID)
	}
	fri_ship.Status = common.PENDING
	db.Create(&fri_ship)
}

func GetFriendsList(userid uint) *[]model.User {
	db := pool.GetDB()
	mres := make([]model.FriendShip, 0)
	res := make([]model.User, 0)
	db.Raw("SELECT * FROM friend_ships WHERE user_id = ? OR friend_id = ?;", userid, userid).Scan(&mres)
	for _, v := range mres {
		var friend model.User
		if v.Status != common.ACCEPTED {
			continue
		}
		friendId := v.FriendId
		if friendId == userid {
			friendId = v.UserId
		}
		db.First(&friend, "id = ?", friendId)
		res = append(res, friend)
	}
	return &res
}
