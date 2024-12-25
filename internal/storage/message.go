package storage

import (
	"github.com/google/uuid"
	"time"
)

// 聊天消息结构体
type ChatMessage struct {
	ID          int       `json:"id"`           // 自增ID
	UUID        uuid.UUID `json:"uuid"`         // UUID
	SenderID    int32     `json:"sender_id"`    // 发送者的用户ID
	ReceiverID  int32     `json:"receiver_id"`  // 接收者的用户、群组ID
	Content     string    `json:"content"`      // 消息内容
	ContentType string    `json:"content_type"` // 消息类型（如文本、图片等）
	MessageType string    `json:"message_type"` // 私聊、群聊等
	SentAt      time.Time `json:"sent_at"`      // 发送时间
	Status      string    `json:"status"`       // 消息状态（如已发送、已送达、已读）
	Attachment  string    `json:"attachment"`   // 附件（如果有）
}
