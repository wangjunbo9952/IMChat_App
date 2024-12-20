package repository

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/internal/model"
	"IMChat_App/pkg/common"
	"IMChat_App/pkg/protocol"
	"encoding/json"
	"fmt"
	"strconv"
)

const NULL_ID uint = 0

func GetMsg(msgReq *common.MsgReq) *[]common.TextMsgResp {
	db := pool.GetDB()

	if msgReq.MessageType == common.MESSAGE_TYPE_USER {
		var resp []common.TextMsgResp
		err := db.Raw(`SELECT from_user_id AS SenderId, content, created_at AS timestamp FROM messages WHERE from_user_id IN (?, ?) AND to_user_id IN (?, ?);`, msgReq.UserId, msgReq.TargetID, msgReq.UserId, msgReq.TargetID).Scan(&resp).Error
		if err != nil {
			return nil
		}
		return &resp
	} else {
		return nil
	}
}

func SaveMessage(message protocol.Message) {
	db := pool.GetDB()
	var fromUser model.User
	fid, _ := strconv.Atoi(message.From)
	db.Find(&fromUser, "id = ?", fid)
	if NULL_ID == fromUser.Model.ID {
		return
	}
	var toUserId int32 = 0

	if message.MessageType == common.MESSAGE_TYPE_USER {
		var toUser model.User
		tid, _ := strconv.Atoi(message.To)
		db.Find(&toUser, "id = ?", tid)
		if NULL_ID == toUser.Model.ID {
			return
		}
		toUserId = int32(toUser.Model.ID)
	}

	saveMessage := model.Message{
		FromUserId:  int32(fromUser.Model.ID),
		ToUserId:    toUserId,
		Content:     message.Content,
		ContentType: int16(message.ContentType),
		MessageType: int16(message.MessageType),
		Url:         message.Url,
	}
	db.Save(&saveMessage)
}

func SaveMessageToRds(message []byte, from, to string) {
	rdb := pool.Rds
	ctx := pool.Ctx
	// 使用用户ID作为键的一部分
	key := from + ":to:" + to

	// 将消息推入Redis列表
	err := rdb.LPush(ctx, key, message).Err()
	if err != nil {
		fmt.Println("Failed to store message in Redis: %s", err)
	}

	// 保持列表长度不超过50
	rdb.LTrim(ctx, key, 0, 49)
}

func GetMsgFromRdb(msgReq *common.MsgReq) *[]common.TextMsgResp {
	rdb := pool.Rds
	ctx := pool.Ctx
	key1 := strconv.Itoa(int(msgReq.UserId)) + ":to:" + strconv.Itoa(int(msgReq.TargetID))
	key2 := strconv.Itoa(int(msgReq.TargetID)) + ":to:" + strconv.Itoa(int(msgReq.UserId))

	messages1, err1 := rdb.LRange(ctx, key1, 0, -1).Result()
	messages2, err2 := rdb.LRange(ctx, key2, 0, -1).Result()

	if err1 == nil && err2 == nil {
		var resp []common.TextMsgResp
		for _, msg := range messages1 {
			var a_resp common.TextMsgResp
			var a_msg model.Message
			err1 = json.Unmarshal([]byte(msg), &a_msg)
			a_resp.SenderId = int(a_msg.FromUserId)
			a_resp.Content = a_msg.Content
			a_resp.Timestamp = a_msg.CreatedAt
			resp = append(resp, a_resp)
		}
		for _, msg := range messages2 {
			var a_resp common.TextMsgResp
			var a_msg model.Message
			err2 = json.Unmarshal([]byte(msg), &a_msg)
			a_resp.SenderId = int(a_msg.FromUserId)
			a_resp.Content = a_msg.Content
			a_resp.Timestamp = a_msg.CreatedAt
			resp = append(resp, a_resp)
		}
		return &resp
	} else {
		return nil
	}
}
