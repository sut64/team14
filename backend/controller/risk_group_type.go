package controller

import (
	"net/http"

	"github.com/sut64/team14/entity"
	"github.com/gin-gonic/gin"
)

// POST /riskgroupType
func CreateRiskGroupType(c *gin.Context) {
	var risk_group_type entity.RiskGroupType
	if err := c.ShouldBindJSON(&risk_group_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&risk_group_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": risk_group_type})
}

// GET /riskgroupType/:id
func GetRiskGroupType(c *gin.Context) {
	var risk_group_type []entity.RiskGroupType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM risk_group_types WHERE id = ?", id).Scan(&risk_group_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": risk_group_type})
}

// GET /riskgroupType
func ListRiskGroupType(c *gin.Context) {
	var risk_group_type []entity.RiskGroupType
	if err := entity.DB().Raw("SELECT * FROM risk_group_types").Scan(&risk_group_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": risk_group_type})
}

// DELETE /riskgroupType/:id
func DeleteRiskGroupType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM risk_group_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "risk_group_type not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /riskgroupType
func UpdateRiskGroupType(c *gin.Context) {
	var risk_group_type entity.RiskGroupType
	if err := c.ShouldBindJSON(&risk_group_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", risk_group_type.ID).First(&risk_group_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "risk_group_type not found"})
		return
	}

	if err := entity.DB().Save(&risk_group_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": risk_group_type})
}

