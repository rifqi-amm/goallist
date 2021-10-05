package database

import (
	"goallist/config"
	"goallist/model"
)

func IsValidBasic(u string) bool {
	return true
}

func IsValid(u string, p string) bool {
	var user model.User
	if err := config.DB.Where("email = ?", u).Find(&user).Error; err != nil {
		return false
	}

	return p == user.Password
}