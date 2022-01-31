package entity

import (
	"gorm.io/gorm"
)

type Officer struct {
	gorm.Model
	Name     string
	Tel      string
	Email    string `gorm:"uniqueIndex"`
	Password string

	//Prevention   []Prevention   `gorm:"foreignKey:OfficerID"`
	Appointment  []Appointment  `gorm:"foreignKey:OfficerID"`
	RoomDataList []RoomDataList `gorm:"foreignKey:OfficerID"`
}
