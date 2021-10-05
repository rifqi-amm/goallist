package model

type Users struct {
	ID       uint   `gorm:"primarykey"`
	Username string `json: "username"`
	Password string `json: "password"`
}