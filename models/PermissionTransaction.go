package models

import (
	"github.com/jinzhu/gorm"
)

type PermissionTransaction struct {
	gorm.Model
	Point              int    `json:"point" form:"point"`
	IdPermissionMaster string `json:"id_permission_master" form:"id_permission_master"`
	// PermissionMaster   PermissionMaster `gorm:"foreignKey:IdPermissionMaster" json:"permission_master" form:"permission_master"`
}

func (Articles) TableName() string {
	return "tbl_permission_transaction"
}
