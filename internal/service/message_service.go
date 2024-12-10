package service

import (
	"IMChat_App/internal/repository"
	"IMChat_App/pkg/common"
)

func GetMsg(msgReq *common.MsgReq) *[]common.TextMsgResp {
	return repository.GetMsg(msgReq)
}
