package database

import (
	"goallist/config"
	"goallist/model"
)

func GetLongterm() []model.Longterm {
	var term []model.Longterm
	config.DB.Preload("Status").Find(&term)
	return term
}

func CreateLongterm(term model.Longterm) model.Longterm {
	config.DB.Preload("Status").Create(&term)
	return term
}

func DeleteLongterm(id string) {
	var term model.Longterm
	config.DB.Where("id = ?", id).Delete(&term)
}

func UpdateLongterm(id string, term model.Longterm) {
	config.DB.Where("id = ?", id).Updates(&term)
}

