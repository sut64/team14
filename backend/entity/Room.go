package entity

import (

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomNumber		string
	Screening		[]Screening `gorm:"foreignKey:RoomID"`
}

