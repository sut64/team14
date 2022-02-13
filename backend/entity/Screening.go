package entity

import (
	"time"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Screening	struct {
	gorm.Model
	BloodPressure		int	
	CongenitalDisease	string	`valid:"required~CongenitalDisease can not be blank"`
	Time			time.Time `valid:"present~Screening Date must be in Present"`
	PatientID		*uint
	Patient			Patient `gorm:"references:id" valid:"-"`
	RoomID			*uint
	Room			Room `gorm:"references:id" valid:"-"`
	SymptomID		*uint
	Symptom			Symptom `gorm:"references:id" valid:"-"`
	OfficerID 		*uint
	Officer       		Officer `gorm:"references:id" valid:"-"`
}
func init() {
	govalidator.CustomTypeTagMap.Set("present", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Before(time.Now()) || t.Equal(time.Now())
	})
}
