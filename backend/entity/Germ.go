package entity

import (
	
	"gorm.io/gorm"
)

type Germ struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`

	Contagious []Contagious `gorm:"foreignKey:GermID"`
}

