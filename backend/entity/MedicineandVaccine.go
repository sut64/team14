package entity

import (
	"time"

	"gorm.io/gorm"
)

type MedicineandVaccine struct {
	gorm.Model
	RegNo  string	`valid:"matches(^[A-Z]{1}\\d{4}$)"`
	Name   string	`valid:"required~Name cannot be blank"`
	MinAge int	`valid:"positive~MinAge does not validate as positive"`
	MaxAge int	`valid:"positive~MaxAge does not validate as positive"`
	Date   time.Time `valid:"notpast~Date must not be past"`

	//DosageForm ทำหน้าที่เป็น FK
	DosageFormID *uint
	DosageForm   DosageForm `gorm:"references:id"`

	//Contagious ทำหน้าที่เป็น FK
	ContagiousID *uint
	Contagious   Contagious `gorm:"references:id"`

	//Age ทำหน้าที่เป็น FK
	AgeID *uint
	Age   Age `gorm:"references:id"`

	//Category ทำหน้าที่เป็น FK
	CategoryID *uint
	Category   Category `gorm:"references:id"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("notpast", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().Add(time.Minute * -5))
	})
	govalidator.CustomTypeTagMap.Set("positive", func(i interface{}, context interface{}) bool {
		num := i
		return num.(int) >= 1
	})
}

