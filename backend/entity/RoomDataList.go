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
	Patient   Patient

	// OfficerID ทำหน้าที่เป็น FK
	OfficerID *uint
	Officer   Officer

	// SpecialistID ทำหน้าที่เป็น FK
	SpecialistID *uint
	Specialist   Specialist

	// RoomID ทำหน้าที่เป็น FK
	RoomID *uint
	Room   RoomDetail
}
