package database

import (
	"goallist/config"
	"goallist/model"
)

func GetShorttermbydate() []model.Shorttermbydate {
	var term []model.Shorttermbydate
	config.DB.Preload("Status").Find(&term)
	return term
}

func CreateShorttermbydate(term model.Shorttermbydate) model.Shorttermbydate {
	config.DB.Create(&term)
	return term
}

func DeleteShorttermbydate(id string) {
	var term model.Shorttermbydate
	config.DB.Where("id = ?", id).Delete(&term)
}

func UpdateShorttermbydate(id string, term model.Shorttermbydate) {
	config.DB.Where("id = ?", id).Updates(&term)
}

