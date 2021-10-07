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

// -------------------- REQUEST -------------------------
func GetUser() []model.User {
	var datauser []model.User
	config.DB.Preload("Status").Find(&datauser)
	return datauser
}

func GetUserByID(id string) model.User {
	var datauser model.User
	config.DB.Preload("Status").Where("id = ?", id).Find(&datauser)
	return datauser
}


func CreateUser(userdata model.User) model.User {
	config.DB.Create(&userdata)
	return userdata
}

func DeleteUser(id string) {
	var datauser model.User
	config.DB.Where("id = ?", id).Delete(&datauser)
}

func UpdateUser(id string, datauser model.User) {
	config.DB.Where("id = ?", id).Updates(&datauser)
}