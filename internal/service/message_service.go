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

func SaveMessageToRds(message []byte, from, to string) {
	repository.SaveMessageToRds(message, from, to)
}

func GetMsgFromRdb(msgReq *common.MsgReq) *[]common.TextMsgResp {
	return repository.GetMsgFromRdb(msgReq)
}
