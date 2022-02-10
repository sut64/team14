package entity

import (
	"time"

	"gorm.io/gorm"
)

type Contagious struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be blank"`
	Symptom    string `valid:"minstringlength(10)~Symptom must be more than 10"`
	Incubation int    `valid:"range(1|90)~Incubation must be between 1-90"`
	Date       time.Time

	GermID *uint
	Germ   Germ `gorm:"references:id" valid:"-"`

	CatchingTypeID *uint
	CatchingType   CatchingType `gorm:"references:id" valid:"-"`

	RiskGroupTypeID *uint
	RiskGroupType   RiskGroupType `gorm:"references:id" valid:"-"`

//	Prevention         []Prevention         `gorm:"foreignKey:ContagiousID"`
	MedicineandVaccine []MedicineandVaccine `gorm:"foreignKey:ContagiousID"`
}

