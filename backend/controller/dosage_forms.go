package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nukanda/project/entity"
)

// POST /dosage_forms
func CreateDosageForm(c *gin.Context) {
	var dosageform entity.DosageForm
	if err := c.ShouldBindJSON(&dosageform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&dosageform).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dosageform})
}

// GET /dosageform/:id
func GetDosageForm(c *gin.Context) {
	var dosageform entity.DosageForm
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM dosage_forms WHERE id = ?", id).Scan(&dosageform).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dosageform})
}

// GET /dosage_forms
func ListDosageForms(c *gin.Context) {
	var dosageforms []entity.DosageForm
	if err := entity.DB().Raw("SELECT * FROM dosage_forms").Scan(&dosageforms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dosageforms})
}

// DELETE /dosage_forms/:id
func DeleteDosageForm(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM dosage_forms WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dosageform not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /dosage_forms
func UpdateDosageForm(c *gin.Context) {
	var dosageform entity.DosageForm
	if err := c.ShouldBindJSON(&dosageform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", dosageform.ID).First(&dosageform); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dosageform not found"})
		return
	}

	if err := entity.DB().Save(&dosageform).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dosageform})
}

