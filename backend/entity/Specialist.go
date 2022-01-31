package entity

import (
	"gorm.io/gorm"
)

type Specialist struct {
	gorm.Model
	Specialist string
	Name       string
	Tel        string
	Email      string
	//Prevention   []Prevention   `gorm:"foreignKey:SpecialistID"`
	Appointment  []Appointment  `gorm:"foreignKey:SpecialistID"`
	RoomDataList []RoomDataList `gorm:"foreignKey:SpecialistID"`
}
