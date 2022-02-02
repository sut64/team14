package entity

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	AppointDate time.Time
	IssueDate   time.Time
	Note        string `valid:"required~Note can not be blank"`
	Number      int

	OfficerID *uint
	Officer   Officer `gorm:"references:id" valid:"-"`

	SpecialistID *uint
	Specialist   Specialist `gorm:"references:id" valid:"-"`

	PatientID *uint
	Patient   Patient `gorm:"references:id" valid:"-"`
}
