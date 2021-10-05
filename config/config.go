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
	DB.AutoMigrate(
		&model.Status{},
		&model.Shortterm{},
		&model.Shorttermbydate{},
		&model.Mediumterm{},
		&model.Mediumtermbydate{},
		&model.Longterm{},
		&model.Longtermbyage{},
		&model.Longtermnolimit{},
	)
}