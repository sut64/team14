package entity

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name          string `valid:"required~Name cannot be blank"`
	Age           uint
	Gender        string
	BloodPressure uint

	Screening    []Screening    `gorm:"foreignKey:PatientID"`
	RoomDataList []RoomDataList `gorm:"foreignKey:PatientID"`
}
