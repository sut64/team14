package entity

import (
	"time"

	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

type Contagious struct {
	gorm.Model
	Name       string    `valid:"required~Name cannot be blank"`
	Symptom    string    `valid:"minstringlength(10)~Symptom must be more than 10"`
	Incubation int       `valid:"range(1|90)~Incubation must be between 1-90"`
	Date       time.Time `valid:"notFuture~Date cannot be in the future"`

	GermID *uint
	Germ   Germ `gorm:"references:id" valid:"-"`

	CatchingTypeID *uint
	CatchingType   CatchingType `gorm:"references:id" valid:"-"`

	RiskGroupTypeID *uint
	RiskGroupType   RiskGroupType `gorm:"references:id" valid:"-"`

//	Prevention         []Prevention         `gorm:"foreignKey:ContagiousID"`
	MedicineandVaccine []MedicineandVaccine `gorm:"foreignKey:ContagiousID"`
}

// ตรวจสอบวันที่ต้องไม่เป็น อนาคต
func init() {
	govalidator.CustomTypeTagMap.Set("notFuture", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return t.Before(now) || t.Equal(now)
	})
}

