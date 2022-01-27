package entity

import (
	"gorm.io/gorm"
)

type DosageForm struct {
	gorm.Model
	DosageForm         string
	MedicineandVaccine []MedicineandVaccine `gorm:"foreignKey:DosageFormID"`
}

