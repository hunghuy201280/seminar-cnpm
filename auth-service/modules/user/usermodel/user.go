package usermodel

import "auth-service/common"

type User struct {
	common.SQLModel
	Phone          string   `json:"phone" gorm:"column:phone;"`
	FacebookId     string   `json:"fb_id" gorm:"column:fb_id;"`
	GoogleId       string   `json:"gg_id" gorm:"column:gg_id;"`
	Password       string   `json:"-" gorm:"column:password;"`
	Salt           string   `json:"-" gorm:"column:salt;"`
	LastName       string   `json:"last_name" gorm:"column:last_name;"`
	FirstName      string   `json:"first_name" gorm:"column:first_name;"`
	Role           UserRole `json:"role" gorm:"column:role;type:ENUM('buyer','seller','agency','admin')"`
	AvatarUrl      string   `json:"avatar_url,omitempty" gorm:"column:avatar_url;"`
	IdentityNumber string   `json:"identity_number" gorm:"column:identity_number;"`
}

func (u User) TableName() string {
	return "users"
}
