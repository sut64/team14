package entity

import (
	"time"

	"gorm.io/gorm")
	
type CatchingType struct {
	gorm.Model
	Title string `gorm:"uniqueIndex"`

	Contagious []Contagious `gorm:"foreignKey:CatchingTypeID"`
}

