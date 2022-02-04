package controller

import (
	"net/http"

	"github.com/sut64/team14/entity"
	"github.com/gin-gonic/gin"

	"github.com/asaskevich/govalidator"
)

// POST /contagious
func CreateContagious(c *gin.Context) {

	var contagious entity.Contagious
	var germ entity.Germ
	var catching_type entity.CatchingType
	var risk_group_type entity.RiskGroupType

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร contagious
	if err := c.ShouldBindJSON(&contagious); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา germ ด้วย id
	if tx := entity.DB().Where("id = ?", contagious.GermID).First(&germ); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "germ not found"})
		return
	}

	// 10: ค้นหา catchingType ด้วย id
	if tx := entity.DB().Where("id = ?", contagious.CatchingTypeID).First(&catching_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "catching_type not found"})
		return
	}

	// 11: ค้นหา riskgroupType ด้วย id
	if tx := entity.DB().Where("id = ?", contagious.RiskGroupTypeID).First(&risk_group_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "risk_group_type not found"})
		return
	}

	// 12: แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(contagious); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: สร้าง contagious
	s := entity.Contagious{
		Germ:          germ,                  // โยงความสัมพันธ์กับ Entity Germ
		CatchingType:  catching_type,         // โยงความสัมพันธ์กับ Entity CatchingType
		RiskGroupType: risk_group_type,       // โยงความสัมพันธ์กับ Entity RiskGroupType
		Name:          contagious.Name,       // ตั้งค่าฟิลด์ name
		Symptom:       contagious.Symptom,    // ตั้งค่าฟิลด์ symptom
		Incubation:    contagious.Incubation, // ตั้งค่าฟิลด์ incubation
		Date:          contagious.Date,       // ตั้งค่าฟิลด์ date

	}

	// 14: บันทึก
	if err := entity.DB().Create(&s).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": s})
}

// GET /contagious/:id
func GetContagious(c *gin.Context) {
	var contagious entity.Contagious
	id := c.Param("id")
	if err := entity.DB().Preload("Germ").Preload("CatchingType").Preload("RiskGroupType").Raw("SELECT * FROM contagious WHERE id = ?", id).Find(&contagious).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": contagious})
}

// GET /contagious
func ListContagious(c *gin.Context) {
	var contagious []entity.Contagious
	if err := entity.DB().Preload("Germ").Preload("CatchingType").Preload("RiskGroupType").Raw("SELECT * FROM contagious").Find(&contagious).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contagious})
}

// DELETE /contagious/:id
func DeleteContagious(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM contagious WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "contagious not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /contagious
func UpdateContagious(c *gin.Context) {
	var contagious entity.Contagious
	if err := c.ShouldBindJSON(&contagious); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", contagious.ID).First(&contagious); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "contagious not found"})
		return
	}

	if err := entity.DB().Save(&contagious).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contagious})
}

