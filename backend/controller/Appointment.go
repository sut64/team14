package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/entity"
)

//post /Appointment

func CreateAppoint(c *gin.Context) {

	var appointment entity.Appointment
	var officer entity.Officer
	var specialist entity.Specialist
	var patient entity.Patient

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา officer ด้วย id
	if tx := entity.DB().Where("id = ?", appointment.OfficerID).First(&officer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
	}

	// ค้นหา specailist ด้วย id
	if tx := entity.DB().Where("id = ?", appointment.SpecialistID).First(&specialist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Specailist not found"})
	}

	// ค้นหา Patient ด้วย id
	if tx := entity.DB().Where("id = ?", appointment.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient not found"})
	}

	AP := entity.Appointment{
		Officer:     officer,
		Specialist:  specialist,
		Patient:     patient,
		AppointDate: appointment.AppointDate,
		Issue_Date:  appointment.Issue_Date,
		Note:        appointment.Note,
		Number:      appointment.Number,
	}

	if err := entity.DB().Create(&AP).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": AP})
}

//GET /:id
func GetAppointment(c *gin.Context) {
	var appointment entity.Appointment
	id := c.Param("id")
	if err := entity.DB().Preload("Specailist").Preload("Officer").Preload("Pateint").Raw(
		"SELECT * FROM Appointment WHERE id = ?", id).Find(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// GET /get list
func ListAppointment(c *gin.Context) {
	var appointment []entity.Appointment
	if err := entity.DB().Preload("Specailist").Preload("Officer").Preload("Pateint").Raw(
		"SELECT * FROM Appointment").Find(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// DELETE /:id
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Appointment WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

//patch
func UpdateAppointment(c *gin.Context) {
	var appointment entity.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", appointment.ID).First(&appointment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Appointment not found"})
		return
	}

	if err := entity.DB().Save(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appointment})
}
