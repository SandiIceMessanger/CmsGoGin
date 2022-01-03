package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	title   string `json:"token" form:"token"`
	content string `json:"id_user" form:"id_user"`
}

func (Articles) TableName() string {
	return "tbl_post"
}
