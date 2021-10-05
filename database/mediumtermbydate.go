package database

import (
	"goallist/config"
	"goallist/model"
)

func GetMediumtermbydate() []model.Mediumtermbydate {
	var term []model.Mediumtermbydate
	config.DB.Preload("Status").Find(&term)
	return term
}

func GetMediumtermbydateByID(id string) model.Mediumtermbydate {
	var term model.Mediumtermbydate
	config.DB.Preload("Status").Where("id = ?", id).Find(&term)
	return term
}

func CreateMediumtermbydate(term model.Mediumtermbydate) model.Mediumtermbydate {
	config.DB.Create(&term)
	return term
}

func DeleteMediumtermbydate(id string) {
	var term model.Mediumtermbydate
	config.DB.Where("id = ?", id).Delete(&term)
}

func UpdateMediumtermbydate(id string, term model.Mediumtermbydate) {
	config.DB.Where("id = ?", id).Updates(&term)
}

