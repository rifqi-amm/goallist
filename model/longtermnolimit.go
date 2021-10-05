package model

type Longtermnolimit struct {
	ID          uint   `gorm:"primarykey"`
	Goals       string `json: "goals" gorm:"not null"`
	Description string `json: "description" gorm:"not null"`
	CreatedAt   int64  `gorm:"autoCreateTime"`
	StatusID    uint   `json: "statusid"`
	Status      Status
}