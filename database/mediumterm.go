package database

import (
	"goallist/config"
	"goallist/model"
)

func GetMediumterm() []model.Mediumterm {
	var term []model.Mediumterm
	config.DB.Preload("Status").Find(&term)
	return term
}

func GetMediumtermByID(id string) model.Mediumterm {
	var term model.Mediumterm
	config.DB.Preload("Status").Where("id = ?", id).Find(&term)
	return term
}

func CreateMediumterm(term model.Mediumterm) model.Mediumterm {
	config.DB.Create(&term)
	return term
}

func DeleteMediumterm(id string) {
	var term model.Mediumterm
	config.DB.Where("id = ?", id).Delete(&term)
}

func UpdateMediumterm(id string, term model.Mediumterm) {
	config.DB.Where("id = ?", id).Updates(&term)
}

