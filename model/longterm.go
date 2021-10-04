package model

type Longterm struct {
	ID        uint   `gorm:"primarykey"`
	Goals     string `json: "goals" gorm:"not null"`
	Deadline  uint   `json: "deadline" gorm:"not null"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	StatusID  uint   `json: "stausid"`
	Status    Status
}