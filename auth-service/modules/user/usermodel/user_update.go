package usermodel

import "auth-service/common"

type UserUpdate struct {
	common.SQLModel
	Phone          *string   `json:"phone" gorm:"column:phone;"`
	LastName       *string   `json:"last_name" gorm:"column:last_name;"`
	FirstName      *string   `json:"first_name" gorm:"column:first_name;"`
	Role           *UserRole `json:"role" gorm:"column:role;type:ENUM('buyer','seller','agency','admin')"`
	AvatarUrl      *string   `json:"avatar_url,omitempty" gorm:"column:avatar_url;"`
	IdentityNumber *string   `json:"identity_number" gorm:"column:identity_number;"`
}

func (u UserUpdate) TableName() string {
	return "users"
}
