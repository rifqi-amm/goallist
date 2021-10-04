package database

import (
	"goallist/config"
	"goallist/model"
)

func GetShortterm() []model.Shortterm {
	var term []model.Shortterm
	config.DB.Preload("Status").Find(&term)
	return term
}

func CreateShortterm(term model.Shortterm) model.Shortterm {
	config.DB.Create(&term)
	return term
}

func DeleteShortterm(id string) {
	var term model.Shortterm
	config.DB.Where("id = ?", id).Delete(&term)
}

func UpdateShortterm(id string, term model.Shortterm) {
	config.DB.Where("id = ?", id).Updates(&term)
}

