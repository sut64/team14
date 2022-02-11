package entity

import (
	"time"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Screening	struct {
	gorm.Model
	Time			time.Time
	PatientID		*uint
	Patient			Patient 
	RoomID	*uint
	Room		Room 
	SymptomID		*uint
	Symptom			Symptom 
	OfficerID 			*uint
	Officer       		Officer
}
