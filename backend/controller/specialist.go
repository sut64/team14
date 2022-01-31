package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/entity"
)

// POST /specialists
func CreateSpecialist(c *gin.Context) {
	var specialist entity.Patient
	if err := c.ShouldBindJSON(&specialist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&specialist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": specialist})
}

// GET /specialist/:id
func GetSpecialist(c *gin.Context) {
	var specialist entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM specialists WHERE id = ?", id).Scan(&specialist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": specialist})
}

// GET /specialists
func ListSpecialists(c *gin.Context) {
	var users []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM specialists").Scan(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// DELETE /specialists/:id
func DeleteSpecialist(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM specialists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "specialist not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /specialists
func UpdateSpecialist(c *gin.Context) {
	var specialist entity.Patient
	if err := c.ShouldBindJSON(&specialist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", specialist.ID).First(&specialist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&specialist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": specialist})
}
