package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type RoomDataList struct {
	gorm.Model
	Day           int       //`valid:"required~Amount of day must more then 0"`
	Note          string    `valid:"required~Note can not be blank"`
	EnterRoomTime time.Time `valid:"past~EnterRoomTime must be in the past"`

	// PatientID ทำหน้าที่เป็น FK
	PatientID *uint
	Patient   Patient `gorm:"references:id" valid:"-"`

	// OfficerID ทำหน้าที่เป็น FK
	OfficerID *uint
	Officer   Officer `gorm:"references:id" valid:"-"`

	// SpecialistID ทำหน้าที่เป็น FK
	SpecialistID *uint
	Specialist   Specialist `gorm:"references:id" valid:"-"`

	// RoomID ทำหน้าที่เป็น FK
	RoomID *uint
	Room   RoomDetail `gorm:"references:id" valid:"-"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Before(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now())
	})
	govalidator.CustomTypeTagMap.Set("present", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Equal(time.Now())
	})
}
