package common

import "time"

type TextMsgResp struct {
	SenderId  int       `json:"senderId"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
