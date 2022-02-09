package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/entity"
)

// POST /medicineand_vaccines
func CreateMedicineandVaccine(c *gin.Context) {

	var medicineandvaccine entity.MedicineandVaccine
	var category entity.Category
	var dosageForm entity.DosageForm
	var age entity.Age
	var contagious entity.Contagious

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร MedicineandVaccine
	if err := c.ShouldBindJSON(&medicineandvaccine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา category ด้วย id
	if tx := entity.DB().Where("id = ?", medicineandvaccine.CategoryID).First(&category); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found"})
		return
	}

	// 11: ค้นหา age ด้วย id
	if tx := entity.DB().Where("id = ?", medicineandvaccine.AgeID).First(&age); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "age not found"})
		return
	}

	// 12: ค้นหา dosageForm ด้วย id
	if tx := entity.DB().Where("id = ?", medicineandvaccine.DosageFormID).First(&dosageForm); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dosageForm not found"})
		return
	}

	// 13: ค้นหา contagious ด้วย id
	if tx := entity.DB().Where("id = ?", medicineandvaccine.ContagiousID).First(&contagious); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "contagious not found"})
		return
	}

	// 14: สร้าง MedicineandVaccine
	MV := entity.MedicineandVaccine{
		Category:   category,                  // โยงความสัมพันธ์กับ Entity category
		Age:        age,                       // โยงความสัมพันธ์กับ Entity age
		DosageForm: dosageForm,                // โยงความสัมพันธ์กับ Entity DosageForm
		Contagious:  contagious,                 // โยงความสัมพันธ์กับ Entity contagious
		RegNo:      medicineandvaccine.RegNo,  // ตั้งค่าฟิลด์ RegNo
		Name:       medicineandvaccine.Name,   // ตั้งค่าฟิลด์ Name
		MinAge:     medicineandvaccine.MinAge, // ตั้งค่าฟิลด์ MinAge
		MaxAge:     medicineandvaccine.MaxAge, // ตั้งค่าฟิลด์ MaxAge
		Date:       medicineandvaccine.Date,   // ตั้งค่าฟิลด์ Date
	}
	
	if _, err := govalidator.ValidateStruct(MV); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 15: บันทึก
	if err := entity.DB().Create(&MV).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": MV})
}

// GET /medicineandvaccine/:id
func GetMedicineandVaccine(c *gin.Context) {
	var medicineandvaccine entity.MedicineandVaccine
	id := c.Param("id")
	if err := entity.DB().Preload("Category").Preload("Age").Preload("DosageForm").Preload("Contagious").Raw("SELECT * FROM medicineand_vaccines WHERE id = ?", id).Find(&medicineandvaccine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicineandvaccine})
}

// GET /medicinesandvaccine
func ListMedicineandVaccines(c *gin.Context) {
	var medicineandvaccines []entity.MedicineandVaccine
	if err := entity.DB().Preload("Category").Preload("Age").Preload("DosageForm").Preload("Contagious").Raw("SELECT * FROM medicineand_vaccines").Find(&medicineandvaccines).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicineandvaccines})
}

// DELETE /medicineandvaccine/:id
func DeleteMedicineandVaccine(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medicineand_vaccines WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicineandVaccine not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /medicinesandvaccine
func UpdateMedicineandVaccine(c *gin.Context) {
	var medicineandvaccine entity.MedicineandVaccine
	if err := c.ShouldBindJSON(&medicineandvaccine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", medicineandvaccine.ID).First(&medicineandvaccine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicineandVaccine not found"})
		return
	}

	if err := entity.DB().Save(&medicineandvaccine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicineandvaccine})
}



