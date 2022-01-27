package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Category           string
	MedicineandVaccine []MedicineandVaccine `gorm:"foreignKey:CategoryID"`
}

