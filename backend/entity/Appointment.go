package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	AppointDate time.Time `valid:"IsFuture~Appointment Date must be in future"`
	IssueDate   time.Time `valid:"-"`
	Note        string    `valid:"required~Note can not be blank"`
	Number      int

	OfficerID *uint
	Officer   Officer `gorm:"references:id" valid:"-"`

	SpecialistID *uint
	Specialist   Specialist `gorm:"references:id" valid:"-"`

	PatientID *uint
	Patient   Patient `gorm:"references:id" valid:"-"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("IsFuture", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now())
	})
}
