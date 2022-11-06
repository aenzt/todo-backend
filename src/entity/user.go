package entity

import (
	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"primary_key" json:"id"`
	CreatedAt int64          `json:"createdAt"`
	UpdatedAt int64          `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy *int64         `json:"createdBy" gorm:"default:null"`
	UpdatedBy *int64         `json:"updatedBy" gorm:"default:null"`
	DeletedBy *int64         `json:"deletedBy" gorm:"default:null"`

	Username string  `json:"username" gorm:"not null;unique;type:varchar(255)"`
	Password string  `json:"-" gorm:"not null;type:text"`
	Name     string  `json:"name" gorm:"not null;type:varchar(255)"`
	Position *string `json:"position" gorm:"type:varchar(255);default:null"`
	IsMale   bool    `json:"isMale" gorm:"not null"`
	RoleID   int64   `json:"roleId" gorm:"index;not null"`
	RoleName string  `json:"roleName" gorm:"-"`
}

type UserProfile struct {
	User     `json:"user"`
	RoleName string `json:"roleName"`
}

type UserParam struct {
	ID       int64  `uri:"id" param:"id"`
	Username string `json:"-" param:"username"`
	PaginationParam
}

type UserLoginInputParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
