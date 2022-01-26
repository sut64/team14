package entity

import (
	"gorm.io/gorm"
)

type RoomDetail struct {
	gorm.Model
	Name string
	Size string
	Cost string

	// 1 roomdetail `เป็นเจ้าของได้หลาย RoomDataList`
	//RoomDataLists []RoomDataList `gorm:"foreignKey:RoomID"`
}
