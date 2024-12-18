package service

import (
	"IMChat_App/internal/repository"
	"IMChat_App/pkg/common"
	"IMChat_App/pkg/protocol"
)

func GetMsg(msgReq *common.MsgReq) *[]common.TextMsgResp {
	return repository.GetMsg(msgReq)
}

func SaveMessage(message *protocol.Message) {
	repository.SaveMessage(*message)
}
