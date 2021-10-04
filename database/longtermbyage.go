package database

import (
	"goallist/config"
	"goallist/model"
)

func GetLongtermbyage() []model.Longtermbyage {
	var term []model.Longtermbyage
	config.DB.Preload("Status").Find(&term)
	return term
}

func CreateLongtermbyage(term model.Longtermbyage) model.Longtermbyage {
	config.DB.Create(&term)
	return term
}

func DeleteLongtermbyage(id string) {
	var term model.Longtermbyage
	config.DB.Where("id = ?", id).Delete(&term)
}

func UpdateLongtermbyage(id string, term model.Longtermbyage) {
	config.DB.Where("id = ?", id).Updates(&term)
}

