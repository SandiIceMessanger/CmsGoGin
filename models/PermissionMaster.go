package models

import (
	"github.com/jinzhu/gorm"
)

type PermissionMaster struct {
	gorm.Model
	Permission string `json:"permission" form:"permission"`
	IdUser     string `json:"id_user" form:"id_user"`
	// User       User   `gorm:"foreignKey:IdUser" json:"user" form:"user"`
}

func (Articles) TableName() string {
	return "tbl_permission_master"
}
