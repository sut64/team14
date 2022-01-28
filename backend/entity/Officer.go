package entity

import (
	"gorm.io/gorm"
)

type Officer struct {
	gorm.Model
	Officer      string
	Name         string
	Tel          string
	Email        string
	Password     string
	Prevention   []Prevention   `gorm:"foreignKey:OfficerID"`
	Appointment  []Appointment  `gorm:"foreignKey:OfficerID"`
	RoomDataList []RoomDataList `gorm:"foreignKey:OfficerID"`
}
