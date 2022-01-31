package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/entity"
)

// POST /ages
func CreateAge(c *gin.Context) {
	var age entity.Age
	if err := c.ShouldBindJSON(&age); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&age).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": age})
}

// GET /age/:id
func GetAge(c *gin.Context) {
	var age entity.Age
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM ages WHERE id = ?", id).Scan(&age).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": age})
}

// GET /ages
func ListAges(c *gin.Context) {
	var ages []entity.Age
	if err := entity.DB().Raw("SELECT * FROM ages").Scan(&ages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ages})
}

// DELETE /ages/:id
func DeleteAge(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM ages WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ages not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /ages
func UpdateAge(c *gin.Context) {
	var age entity.Age
	if err := c.ShouldBindJSON(&age); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", age.ID).First(&age); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ages not found"})
		return
	}

	if err := entity.DB().Save(&age).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": age})
}
