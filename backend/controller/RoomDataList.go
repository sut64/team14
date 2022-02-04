package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/sut64/team14/entity"

	"github.com/gin-gonic/gin"
)

func CreateRoomDataList(c *gin.Context) {

	var officer entity.Officer
	var patient entity.Patient
	var specialist entity.Specialist
	var roomdetail entity.RoomDetail
	var roomdatalist entity.RoomDataList

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร roomdatalist
	if err := c.ShouldBindJSON(&roomdatalist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//9: ค้นหา officer ด้วย id
	if tx := entity.DB().Where("id = ?", roomdatalist.OfficerID).First(&officer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
		return
	}

	// 10: ค้นหา patient ด้วย id
	if tx := entity.DB().Where("id = ?", roomdatalist.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	// 11: ค้นหา specialist ด้วย id
	if tx := entity.DB().Where("id = ?", roomdatalist.SpecialistID).First(&specialist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "specialist not found"})
		return
	}

	// 12: ค้นหา roomdetail ด้วย id
	if tx := entity.DB().Where("id = ?", roomdatalist.RoomID).First(&roomdetail); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomdetail not found"})
		return
	}

	// 13: สร้าง roomdatalist
	rl := entity.RoomDataList{
		Officer:       officer,    // โยงความสัมพันธ์กับ Entity Officer
		Patient:       patient,    // โยงความสัมพันธ์กับ Entity Patient
		Specialist:    specialist, // โยงความสัมพันธ์กับ Entity Specialist
		Room:          roomdetail, // โยงความสัมพันธ์กับ Entity Roomdetail
		Day:           roomdatalist.Day,
		Note:          roomdatalist.Note,
		EnterRoomTime: roomdatalist.EnterRoomTime, // ตั้งค่าฟิลด์ EnterRoomDateTime
	}

	//validation field Number ต้องมีค่าเป็นบวกเท่านั้น
	if value := govalidator.IsPositive(float64(float64(roomdatalist.Day))); !value {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount of day must more then 0"})
		return
	}

	if _, err := govalidator.ValidateStruct(roomdatalist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&rl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rl})
}

// GET /room_data_list/:id
func GetRoomDataList(c *gin.Context) {
	var roomdatalist entity.RoomDataList
	id := c.Param("id")
	if err := entity.DB().Preload("Officer").Preload("Patient").Preload("Specialist").Preload("Room").Raw("SELECT * FROM room_data_lists WHERE id = ?", id).Find(&roomdatalist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomdatalist})
}

// GET /room_data_lists
func ListRoomDataList(c *gin.Context) {
	var roomdatalists []entity.RoomDataList
	if err := entity.DB().Preload("Officer").Preload("Patient").Preload("Specialist").Preload("Room").Raw("SELECT * FROM room_data_lists").Find(&roomdatalists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomdatalists})
}

// DELETE /room_data_lists/:id

func DeleteRoomDataList(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM room_data_lists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_data_lists not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /room_data_lists
func UpdateRoomDataList(c *gin.Context) {
	var roomdatalist entity.RoomDataList
	if err := c.ShouldBindJSON(&roomdatalist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", roomdatalist.ID).First(&roomdatalist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_data_lists not found"})
		return
	}

	if err := entity.DB().Save(&roomdatalist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomdatalist})
}
