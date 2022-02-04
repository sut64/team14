package entity

import (
	"time"

	"gorm.io/gorm"
)

type Contagious struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be blank"`
	Symptom    string
	Incubation int
	Date       time.Time

	GermID *uint
	Germ   Germ `gorm:"references:id" valid:"-"`

	CatchingTypeID *uint
	CatchingType   CatchingType `gorm:"references:id" valid:"-"`

	RiskGroupTypeID *uint
	RiskGroupType   RiskGroupType `gorm:"references:id" valid:"-"`

//	Prevention         []Prevention         `gorm:"foreignKey:ContagiosID"`
	MedicineandVaccine []MedicineandVaccine `gorm:"foreignKey:ContagiuosID"`
}

