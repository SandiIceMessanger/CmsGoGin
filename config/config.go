package config

import (
	"CMSGo/models"
	"fmt"

	"github.com/jinzhu/gorm"

	"gorm.io/driver/mysql"

	_ "github.com/heroku/x/hmetrics/onload"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "sandi1988",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "db_cms_golang_development",
	}

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config["DB_Username"],
			config["DB_Password"],
			config["DB_Host"],
			config["DB_Port"],
			config["DB_Name"],
		)

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.PasswordForget{})
	DB.AutoMigrate(&models.PermissionMaster{})
	DB.AutoMigrate(&models.PermissionTransaction{})
	DB.AutoMigrate(&models.User{})

}
