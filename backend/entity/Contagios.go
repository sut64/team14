package entity

import (
	"gorm.io/gorm"
)

type Contagios struct {
	gorm.Model
	Contagios          string
	Prevention         []Prevention         `gorm:"foreignKey:ContagiosID"`
	MedicineandVaccine []MedicineandVaccine `gorm:"foreignKey:ContagiosID"`
}
