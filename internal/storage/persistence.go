package storage

import (
	"IMChat_App/pkg/common"
	"IMChat_App/pkg/common/constant"
	pro "IMChat_App/pkg/proto"
	"IMChat_App/pkg/utils"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"time"
)

func ConvertMsg(fromMsg *pro.Message, toMsg *ChatMessage) {
	timeStr := "2022-12-25 12:00:00"
	toMsg.Content = fromMsg.Content
	tm, _ := time.Parse(timeStr, fromMsg.SendAt)
	toMsg.SentAt = tm
	toMsg.ContentType = fromMsg.ContentType
	toMsg.MessageType = fromMsg.MessageType
	toMsg.Status = fromMsg.Status
	toMsg.SenderID = fromMsg.Sender
	toMsg.ReceiverID = fromMsg.Receiver
	toMsg.Attachment = fromMsg.Attachment
	toMsg.UUID = utils.GetUuid()
	return
}

func StoreMsg(msg []byte) error {
	db := GetDB()

	var cMsg ChatMessage
	var tmpMsg pro.Message
	err := proto.Unmarshal(msg, &tmpMsg)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return err
	}
	ConvertMsg(&tmpMsg, &cMsg)

	// 使用GORM创建记录
	result := db.Create(&cMsg)

	// 检查是否有错误发生
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Register(user *User) bool {
	db := GetDB()
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

func Login(user *User) bool {
	db := GetDB()

	var recUser User
	db.First(&recUser, "account = ?", user.Account)
	if recUser.Password != user.Password {
		// 账户不存在
		return false
	}
	return true
}

func SearchUser(account string) (bool, *User) {
	db := GetDB()

	var user User
	db.First(&user, "account = ?", account)

	if user.Account != account {
		return false, nil
	}

	return true, &user

}

func AddUser(relation *common.AddUserReq) {
	db := GetDB()
	var fri_ship FriendShip
	if relation.UserID <= relation.FriendID {
		fri_ship.UserId = uint(relation.UserID)
		fri_ship.FriendId = uint(relation.FriendID)
	} else {
		fri_ship.UserId = uint(relation.FriendID)
		fri_ship.FriendId = uint(relation.UserID)
	}
	fri_ship.Status = constant.PENDING
	db.Create(&fri_ship)
}

func GetFriendsList(userid uint) *[]User {
	db := GetDB()
	mres := make([]FriendShip, 0)
	res := make([]User, 0)
	db.Raw("SELECT * FROM friend_ships WHERE user_id = ? OR friend_id = ?;", userid, userid).Scan(&mres)
	for _, v := range mres {
		var friend User
		if v.Status != constant.ACCEPTED {
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
