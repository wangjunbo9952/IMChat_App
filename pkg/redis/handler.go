package redis

import (
	"IMChat_App/pkg/common/constant"
	"IMChat_App/pkg/common/response"
	pro "IMChat_App/pkg/proto"
	"github.com/gogo/protobuf/proto"
	"strconv"
	"time"
)

func UpdateHistory(message []byte) error {
	var msg pro.Message
	err := proto.Unmarshal(message, &msg)
	if err != nil {
		return err
	}

	var from, to string
	if msg.MessageType == constant.MESSAGE_TYPE_USER {
		// Redis只保存一端，保证小ID在前
		if msg.Sender < msg.Receiver {
			from = strconv.Itoa(int(msg.Sender))
			to = strconv.Itoa(int(msg.Receiver))
		} else {
			to = strconv.Itoa(int(msg.Sender))
			from = strconv.Itoa(int(msg.Receiver))
		}
	} else if msg.MessageType == constant.MESSAGE_TYPE_GROUP {
		from = ""
		to = strconv.Itoa(int(msg.Receiver))
	}

	// 使用用户ID作为键的一部分
	key := from + ":to:" + to

	// 将消息推入Redis列表
	err = Rdb.LPush(Ctx, key, message).Err()
	if err != nil {
		return err
	}

	// 保持列表长度不超过50
	Rdb.LTrim(Ctx, key, 0, 49)

	return nil
}

func GetHistoryMsg(key string) *[]response.MsgResp {
	// 从Redis中读取消息
	messages, err := Rdb.LRange(Ctx, key, 0, -1).Result()

	if err == nil {
		var resp []response.MsgResp

		for _, msg := range messages {
			timeStr := "2022-12-25 12:00:00"
			var a_resp response.MsgResp
			var a_msg pro.Message
			err = proto.Unmarshal([]byte(msg), &a_msg)
			tm, _ := time.Parse(timeStr, a_msg.SendAt)
			a_resp.SenderId = a_msg.Sender
			a_resp.Content = a_msg.Content
			a_resp.Timestamp = tm
			a_resp.ContentType = a_msg.ContentType
			resp = append(resp, a_resp)
		}

		return &resp
	} else {
		return nil
	}
}
