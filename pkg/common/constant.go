package common

const (
	// 好友申请状态
	PENDING  = "pending"
	ACCEPTED = "accepted"
	REJECTED = "rejected"
	BLOCKED  = "blocked"

	// 消息类型，单聊或者群聊
	MESSAGE_TYPE_USER  = 1
	MESSAGE_TYPE_GROUP = 2

	// 消息内容类型
	TEXT         = 1
	FILE         = 2
	IMAGE        = 3
	AUDIO        = 4
	VIDEO        = 5
	AUDIO_ONLINE = 6
	VIDEO_ONLINE = 7

	// 心跳监测
	HEAT_BEAT = "heatbeat"
	PONG      = "pong"
)
