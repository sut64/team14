package controller

import (
	"net/http"

	"github.com/sut64/team14/entity"
	"github.com/gin-gonic/gin"
)

// POST /germ
func CreateGerm(c *gin.Context) {
	var germ entity.Germ
	if err := c.ShouldBindJSON(&germ); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&germ).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": germ})
}

// GET /germ/:id
func GetGerm(c *gin.Context) {
	var germ []entity.Germ
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM germs WHERE id = ?", id).Scan(&germ).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": germ})
}

// GET /germ
func ListGerm(c *gin.Context) {
	var germ []entity.Germ
	if err := entity.DB().Raw("SELECT * FROM germs").Scan(&germ).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": germ})
}

// DELETE /germ/:id
func DeleteGerm(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM germs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "germ not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /germ
func UpdateGerm(c *gin.Context) {
	var germ entity.Germ
	if err := c.ShouldBindJSON(&germ); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", germ.ID).First(&germ); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "germ not found"})
		return
	}

	if err := entity.DB().Save(&germ).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": germ})
}

