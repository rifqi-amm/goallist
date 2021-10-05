package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"goallist/model"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goallist_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

}

func InitMigration() {
	DB.AutoMigrate(&model.Users{})
	DB.AutoMigrate(&model.Status{})
	DB.AutoMigrate(&model.Shortterm{})
	DB.AutoMigrate(&model.Shorttermbydate{})
	DB.AutoMigrate(&model.Mediumterm{})
	DB.AutoMigrate(&model.Mediumtermbydate{})
	DB.AutoMigrate(&model.Longterm{})
	DB.AutoMigrate(&model.Longtermbyage{})
	DB.AutoMigrate(&model.Longtermnolimit{})
}