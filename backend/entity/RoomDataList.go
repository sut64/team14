package entity

import (
	"time"

	"gorm.io/gorm"
)

type RoomDataList struct {
	gorm.Model
	Day           int
	Note          string
	EnterRoomTime time.Time

	// PatientID ทำหน้าที่เป็น FK
	PatientID *uint
	Patient   Patient `gorm:"references:id"`

	// OfficerID ทำหน้าที่เป็น FK
	OfficerID *uint
	Officer   Officer `gorm:"references:id"`

	// SpecialistID ทำหน้าที่เป็น FK
	SpecialistID *uint
	Specialist   Specialist `gorm:"references:id"`

	// RoomID ทำหน้าที่เป็น FK
	RoomID *uint
	Room   RoomDetail `gorm:"references:id"`
}
