package controller


import (
		//"time"

        "github.com/sut64/team14/entity"

        "github.com/gin-gonic/gin"

        "net/http"

)

// POST /pat

func CreateScreening(c *gin.Context) {

	var symptom entity.Symptom
	var room entity.Room
	var name entity.Patient
	var screening entity.Screening
	var officers entity.Officer



	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Patient
	if err := c.ShouldBindJSON(&screening); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา symptom ด้วย id
	if tx := entity.DB().Where("id = ?", screening.SymptomID).First(&symptom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Symptom not found"})
		return
	}

	// 9: ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", screening.OfficerID).First(&officers); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Officer not found"})
		return
	}

	// 11: ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", screening.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
		return
	}

	// 12: ค้นหา namr ด้วย id
	if tx := entity.DB().Where("id = ?", screening.PatientID).First(&name); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient not found"})
		return
	}

	// 15: สร้าง Screening
	sc := entity.Screening{
		Time: 			screening.Time,	   		// 14: ดึงเวลาปัจจุบัน
		Symptom:	 	symptom,         // โยงความสัมพันธ์กับ Entity symptom
		Room:       	room,               // โยงความสัมพันธ์กับ Entity room
		Patient: 		name,               // โยงความสัมพันธ์กับ Entity name
		Officer: 		officers,

	}

	// 16: บันทึก
	if err := entity.DB().Create(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sc})

}

//get id

func GetScreening(c *gin.Context) {
	var screening entity.Screening
	id := c.Param("id")
	if err := entity.DB().Preload("Symptom").Preload("Room").Preload("Patient").Preload("Officer").Raw("SELECT * FROM screenings WHERE id = ?", id).Find(&screening).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": screening})
}

// GET /pats


func ListScreening(c *gin.Context) {


	var screening []entity.Screening
	if err := entity.DB().Preload("Symptom").Preload("Room").Preload("Patient").Preload("Officer").Raw("SELECT * FROM screenings").Find(&screening).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": screening})

}

//ScreeningScreen
/*func GetScreening(c *gin.Context) {

	var screening []entity.Screening
	if err := entity.DB().Preload("Patient").Preload("Room").Preload("Symptom")/*.Preload("User").Raw("SELECT * FROM screenings").Find(&screening).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": screening})

}*/

// DELETE /users/:id

func DeleteScreening(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM screenings WHERE id = ?", id); tx.RowsAffected == 0 {

			c.JSON(http.StatusBadRequest, gin.H{"error": "screenings not found"})

			return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}
// PATCH /users

func UpdateScreening(c *gin.Context) {

	var screening entity.Screening

	if err := c.ShouldBindJSON(&screening); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return

	}


	if tx := entity.DB().Where("id = ?", screening.ID).First(&screening); tx.RowsAffected == 0 {

			c.JSON(http.StatusBadRequest, gin.H{"error": "screenings not found"})

			return

	}


	if err := entity.DB().Save(&screening).Error; err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return

	}


	c.JSON(http.StatusOK, gin.H{"data": screening})

}
