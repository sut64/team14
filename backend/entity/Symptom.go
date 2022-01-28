package entity

import (
	"gorm.io/gorm"
)

type Symptom struct {
	gorm.Model
	State			string
	Period			uint
	Screening		[]Screening `gorm:"foreignKey:SymptomID"`
}
