package repository

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/pkg/common"
	"fmt"
)

func GetMsg(msgReq *common.MsgReq) *[]common.TextMsgResp {
	db := pool.GetDB()

	if msgReq.MessageType == common.MESSAGE_TYPE_USER {
		var resp []common.TextMsgResp
		err := db.Raw(`SELECT from_user_id AS SenderId, content, created_at AS timestamp FROM messages WHERE from_user_id IN (?, ?) AND to_user_id IN (?, ?);`, msgReq.UserId, msgReq.TargetID, msgReq.UserId, msgReq.TargetID).Scan(&resp).Error
		if err != nil {
			fmt.Println("Error during Scan:", err)
		}
		return &resp
	} else {
		return nil
	}
}
