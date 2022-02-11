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
	Time			time.Time `valid:"IsPresent~Screening Date must be in Present"`
	PatientID		*uint
	Patient			Patient 
	RoomID	*uint
	Room		Room 
	SymptomID		*uint
	Symptom			Symptom 
	OfficerID 			*uint
	Officer       		Officer
}
func init() {
	govalidator.CustomTypeTagMap.Set("IsFuture", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("IsPresent", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		n := t.Format("2006-January-02")
		return n == time.Now().Format("2006-January-02")
	})
}

