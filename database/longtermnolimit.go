package database

import (
	"goallist/config"
	"goallist/model"
)

func GetLongtermnolimit() []model.Longtermnolimit {
	var term []model.Longtermnolimit
	config.DB.Preload("Status").Find(&term)
	return term
}

func CreateLongtermnolimit(term model.Longtermnolimit) model.Longtermnolimit {
	config.DB.Create(&term)
	return term
}

func DeleteLongtermnolimit(id string) {
	var term model.Longtermnolimit
	config.DB.Where("id = ?", id).Delete(&term)
}

func UpdateLongtermnolimit(id string, term model.Longtermnolimit) {
	config.DB.Where("id = ?", id).Updates(&term)
}

