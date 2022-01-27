package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/entity"
)

// POST /room_details
func CreateRoomDetail(c *gin.Context) {
	var room_detail entity.Patient
	if err := c.ShouldBindJSON(&room_detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&room_detail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room_detail})
}

// GET /room_detail/:id
func GetRoomDetails(c *gin.Context) {
	var room_detail entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM room_details WHERE id = ?", id).Scan(&room_detail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room_detail})
}

// GET /room_details
func ListRoomDetails(c *gin.Context) {
	var users []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM room_details").Scan(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// DELETE /room_details/:id
func DeleteRoomDetail(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM room_details WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_detail not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /room_details
func UpdateRoomDetail(c *gin.Context) {
	var room_detail entity.Patient
	if err := c.ShouldBindJSON(&room_detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", room_detail.ID).First(&room_detail); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&room_detail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room_detail})
}
