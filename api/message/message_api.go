package message

import (
	"IMChat_App/internal/storage"
	"IMChat_App/pkg/common/response"
	"IMChat_App/pkg/redis"
)

// 消息存储的api
func StoreMsg(msg []byte) error {
	err := storage.StoreMsg(msg)
	if err != nil {
		return err
	}
	return nil
}

// 根据日期查询消息的api
func SearchByDate() {

}

// 从数据库中获取更多历史消息的api
func MoreHistoryMsg() {

}

// 更新Redis的历史消息队列
func UpdateHistory(message []byte) error {
	return redis.UpdateHistory(message)
}

// 从Redis中获取历史消息,要求前端拼接好key
// 拼接key规则为：
// 私聊：小用户ID :to: 大用户ID
// 群聊：:to: 群ID
func GetHistoryMsg(key string) *[]response.MsgResp {
	return redis.GetHistoryMsg(key)
}
