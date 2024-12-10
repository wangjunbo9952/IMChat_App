package common

type AddUserReq struct {
	UserID   float64 `json:"userId"`
	FriendID float64 `json:"friendId"`
}

type MsgReq struct {
	MessageType int32 `json:"messageType"`
	UserId      uint  `json:"userId"`
	TargetID    uint  `json:"targetID"`
}
