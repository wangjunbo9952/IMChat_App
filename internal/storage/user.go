package storage

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Uuid      string    `json:"uuid" gorm:"type:varchar(150);not null;unique_index:idx_uuid;comment:'uuid'"`
	Account   string    `json:"account" form:"account" binding:"required" gorm:"unique;not null; comment:'账号'"`
	Password  string    `json:"password" form:"password" binding:"required" gorm:"type:varchar(150);not null; comment:'密码'"`
	Username  string    `json:"username" gorm:"comment:'名称'"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(150);comment:'头像'"`
	Email     string    `json:"email" gorm:"type:varchar(80);column:email;comment:'邮箱'"`
	BirthDate time.Time `json:"birth_date,omitempty" gorm:"default:null"`
	Gender    string    `json:"gender,omitempty"`
	Phone     string    `json:"phone,omitempty"`
}

type FriendShip struct {
	gorm.Model
	UserId   uint   `gorm:"primaryKey;not null" json:"user_id"`
	FriendId uint   `gorm:"primaryKey;not null" json:"friend_id"`
	Status   string `gorm:"type:enum('pending', 'accepted', 'rejected', 'blocked');default:'pending';not null" json:"status"`
}
