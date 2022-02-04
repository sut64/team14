package controller

import (
	"net/http"

	"github.com/sut64/team14/entity"
	"github.com/gin-gonic/gin"
)

// POST /catchingtype
func CreateCatchingType(c *gin.Context) {
	var catching_type entity.CatchingType
	if err := c.ShouldBindJSON(&catching_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&catching_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": catching_type})
}

// GET /catchingtype/:id
func GetCatchingType(c *gin.Context) {
	var catching_type []entity.CatchingType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM catching_types WHERE id = ?", id).Scan(&catching_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": catching_type})
}

// GET /catchingtype
func ListCatchingType(c *gin.Context) {
	var catching_type []entity.CatchingType
	if err := entity.DB().Raw("SELECT * FROM catching_types").Scan(&catching_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": catching_type})
}

// DELETE /catchingtype/:id
func DeleteCatchingType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM catching_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "catching_type not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /catchingtype
func UpdateCatchingType(c *gin.Context) {
	var catching_type entity.CatchingType
	if err := c.ShouldBindJSON(&catching_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", catching_type.ID).First(&catching_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "catching_type not found"})
		return
	}

	if err := entity.DB().Save(&catching_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": catching_type})
}

