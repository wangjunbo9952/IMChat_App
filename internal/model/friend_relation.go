package model

import "gorm.io/gorm"

type FriendShip struct {
	gorm.Model
	UserId   uint   `gorm:"primaryKey;not null" json:"user_id"`
	FriendId uint   `gorm:"primaryKey;not null" json:"friend_id"`
	Status   string `gorm:"type:enum('pending', 'accepted', 'rejected', 'blocked');default:'pending';not null" json:"status"`
}
