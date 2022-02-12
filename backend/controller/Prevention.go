package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/entity"
)

// POST /prevention
func CreatePrevention(c *gin.Context) {
	var contagious entity.Contagious
	var officer entity.Officer
	var specialist entity.Specialist
	var prevention entity.Prevention
	if err := c.ShouldBindJSON(&prevention); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา officer ด้วย id
	if tx := entity.DB().Where("id = ?", prevention.OfficerID).First(&officer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
		return
	}

	// 10: ค้นหา contagious ด้วย id
	if tx := entity.DB().Where("id = ?", prevention.ContagiousID).First(&contagious); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "contagious not found"})
		return
	}

	// 11: ค้นหา specialist ด้วย id
	if tx := entity.DB().Where("id = ?", prevention.SpecialistID).First(&specialist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "specialist not found"})
		return
	}

	// 12: สร้าง Prevention
	p := entity.Prevention{
		Officer:    officer,               // โยงความสัมพันธ์กับ Entity Officer
		Contagious: contagious,            // โยงความสัมพันธ์กับ Entity contagious
		Disease:    prevention.Disease,    // ตั้งค่าฟิลด์ Disease
		Specialist: specialist,            // โยงความสัมพันธ์กับ Entity specialist
		Protection: prevention.Protection, // ตั้งค่าฟิลด์ Protection
		Date:       prevention.Date,       // ตั้งค่าฟิลด์ Date
		Age:        prevention.Age,        // ตั้งค่าฟิลด์ Age
	}

	if _, err := govalidator.ValidateStruct(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": prevention})
}

// GET /prevention/:id
func GetPrevention(c *gin.Context) {
	var prevention entity.Prevention
	id := c.Param("id")
	if err := entity.DB().Preload("Officer").Preload("Contagious").Preload("Specialist").Raw("SELECT * FROM prevention WHERE id = ?", id).Find(&prevention).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prevention})
}

// GET /preventions
func ListPrevention(c *gin.Context) {
	var preventions []entity.Prevention
	if err := entity.DB().Preload("Officer").Preload("Contagious").Preload("Specialist").Raw("SELECT * FROM preventions").Find(&preventions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": preventions})
}

// DELETE /preventions/:id
func DeletePrevention(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM preventions WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "preventions not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /preventions
func UpdatePrevention(c *gin.Context) {
	var prevention entity.Prevention
	if err := c.ShouldBindJSON(&prevention); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", prevention.ID).First(&prevention); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prevention not found"})
		return
	}

	if err := entity.DB().Save(&prevention).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prevention})
}
