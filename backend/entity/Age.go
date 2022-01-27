package entity

import (
	"gorm.io/gorm"
)

type Age struct {
	gorm.Model
	Age                string
	MedicineandVaccine []MedicineandVaccine `gorm:"foreignKey:AgeID"`
}

