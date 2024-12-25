package response

import "time"

type MsgResp struct {
	SenderId    int32     `json:"senderId"`
	Content     string    `json:"content"`
	ContentType string    `json:"contentType"`
	Timestamp   time.Time `json:"timestamp"`
}
