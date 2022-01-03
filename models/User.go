package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name" form:"name"`
	Email     string `gorm:"unique" json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Token     string `json:"token" form:"token"`
	Status    bool   `json:"status" form:"status"`
	TypeUser  string `json:"type_user" form:"type_user"`
	CreatedBy string `json:"created_by" form:"created_by"`
}

func (User) TableName() string {
	return "tbl_user"
}
