package repository

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/internal/model"
	"IMChat_App/pkg/common"
	"IMChat_App/pkg/protocol"
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
