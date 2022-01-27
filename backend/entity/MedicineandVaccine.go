package entity

import (
	"time"

	"gorm.io/gorm"
)

type MedicineandVaccine struct {
	gorm.Model
	RegNo  string
	Name   string
	MinAge uint8
	MaxAge uint8
	Date   time.Time

	//DosageForm ทำหน้าที่เป็น FK
	DosageFormID *uint
	DosageForm   DosageForm `gorm:"references:id"`

	//Contagios ทำหน้าที่เป็น FK
	ContagiosID *uint
	Contagios   Contagios `gorm:"references:id"`

	//Age ทำหน้าที่เป็น FK
	AgeID *uint
	Age   Age `gorm:"references:id"`

	//Category ทำหน้าที่เป็น FK
	CategoryID *uint
	Category   Category `gorm:"references:id"`
}

