package constant

const (
	// 好友申请状态
	PENDING  = "pending"
	ACCEPTED = "accepted"
	REJECTED = "rejected"
	BLOCKED  = "blocked"

	// 消息类型，单聊或者群聊
	MESSAGE_TYPE_USER  = "private_chat"
	MESSAGE_TYPE_GROUP = "public_chat"

	// 消息内容类型
	TEXT  = "text_msg"
	FILE  = "file_msg"
	IMAGE = "image_msg"
	AUDIO = "audio_msg"
	VIDEO = "video_msg"

	// 心跳监测
	HEAT_BEAT = "heatbeat"
	PONG      = "pong"
)
