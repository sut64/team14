package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
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
		IssueDate:   appointment.IssueDate,
		Note:        appointment.Note,
		Number:      appointment.Number,
	}

	//validation field Number ต้องมีค่าเป็นบวกเท่านั้น
	if value := govalidator.IsPositive(float64(float64(appointment.Number))); !value {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount of day must more then 0"})
		return
	}

	if _, err := govalidator.ValidateStruct(appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
	if err := entity.DB().Preload("Specialist").Preload("Officer").Preload("Patient").Raw(
		"SELECT * FROM appointments WHERE id = ?", id).Find(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// GET /get list
func ListAppointment(c *gin.Context) {
	var appointment []entity.Appointment
	if err := entity.DB().Preload("Specialist").Preload("Officer").Preload("Patient").Raw(
		"SELECT * FROM appointments").Find(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// DELETE /:id
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM appointments WHERE id = ?", id); tx.RowsAffected == 0 {
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
