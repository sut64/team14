package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/entity"
	"golang.org/x/crypto/bcrypt"
)

// POST /officers
func CreateOfficer(c *gin.Context) {
	var officer entity.Officer
	if err := c.ShouldBindJSON(&officer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(officer.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	officer.Password = string(bytes)

	if err := entity.DB().Create(&officer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officer})
}

// GET /officer/:id
func GetOfficer(c *gin.Context) {
	var officer entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM officers WHERE id = ?", id).Scan(&officer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officer})
}

// GET /officers
func ListOfficers(c *gin.Context) {
	var users []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM officers").Scan(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// DELETE /officers/:id
func DeleteOfficer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM officers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /officers
func UpdateOfficer(c *gin.Context) {
	var officer entity.Patient
	if err := c.ShouldBindJSON(&officer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", officer.ID).First(&officer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&officer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officer})
}
