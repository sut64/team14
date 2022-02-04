package entity

import (
	"time"

	"gorm.io/gorm"
)


type RiskGroupType struct {
	gorm.Model
	Title string `gorm:"uniqueIndex"`

	Contagious []Contagious `gorm:"foreignKey:RiskGroupTypeID"`
}

