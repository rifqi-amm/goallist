package model

type Status struct {
	ID         uint   `gorm:"primarykey"`
	StatusName string `json: "goals"`
}