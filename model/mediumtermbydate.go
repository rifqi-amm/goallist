package model

import "time"

type Mediumtermbydate struct {
	ID        uint      `gorm:"primarykey"`
	Goals     string    `json: "goals" gorm:"not null"`
	Deadline  time.Time `json: "deadline" gorm:"not null"`
	CreatedAt int64     `gorm:"autoCreateTime"`
	StatusID  uint		`json: "stausid"`
	Status    Status
}