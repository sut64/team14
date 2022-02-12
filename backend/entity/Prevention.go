package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Prevention struct {
	gorm.Model
	Disease    string    `valid:"required~Disease cannot be blank"`
	Protection string    `valid:"minstringlength(5)~Protection must be more than 5"`
	Age        int       `valid:"positive~Age does not validate as positive"`
	Date       time.Time `valid:"notFuture~Date cannot be in the future"`

	OfficerID *uint
	Officer   Officer `gorm:"references:id" valid:"-"`

	ContagiousID *uint
	Contagious   Contagious `gorm:"references:id" valid:"-"`

	SpecialistID *uint
	Specialist   Specialist `gorm:"references:id" valid:"-"`
}

// ตรวจสอบวันที่ต้องไม่เป็น อนาคต
func init() {
	govalidator.CustomTypeTagMap.Set("positive", func(i interface{}, context interface{}) bool {
		num := i
		return num.(int) >= 1
	})
	govalidator.CustomTypeTagMap.Set("notFuture", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return t.Before(now) || t.Equal(now)
	})
}
