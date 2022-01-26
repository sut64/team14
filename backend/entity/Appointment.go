package entity

import (
	"time"

	"gorm.io/gorm"
)

type Officer struct {
	gorm.Model
	Name  string
	Tel   string
	Email string
}

type Specailist struct {
	gorm.Model
	Name  string
	Tel   string
	Email string
}

type Patient struct {
	gorm.Model
	Name    string
	Tel     string
	Email   string
	Symptom string
}

type Appointment struct {
	gorm.Model
	AppointDate time.Time
	Issue_Date  time.Time
	Note        string
	Number      int

	OfficerID *uint
	Officer   Officer

	SpecailistID *uint
	Specailist   Specailist

	PatientID *uint
	Patient   Patient
}
