package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sut64/team14/controller"
	"github.com/sut64/team14/entity"
	"github.com/sut64/team14/middlewares"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())
	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			//Officer Routes
			protected.GET("/officers", controller.ListOfficers)
			protected.GET("/officer/:id", controller.GetOfficer)

			protected.PATCH("/officers", controller.UpdateOfficer)
			protected.DELETE("/officers/:id", controller.DeleteOfficer)

			// Patient Routes
			protected.GET("/patients", controller.ListPatients)
			protected.GET("/patient/:id", controller.GetPatient)
			protected.POST("/patients", controller.CreatePatient)
			protected.PATCH("/patients", controller.UpdatePatient)
			protected.DELETE("/patients/:id", controller.DeletePatient)

			// Specialist Routes
			protected.GET("/specialists", controller.ListSpecialists)
			protected.GET("/specialist/:id", controller.GetSpecialist)
			protected.POST("/specialists", controller.CreateSpecialist)
			protected.PATCH("/specialists", controller.UpdateSpecialist)
			protected.DELETE("/specialists/:id", controller.DeleteSpecialist)

			// RoomDetail Routes
			protected.GET("/room_details", controller.ListRoomDetails)
			protected.GET("/room_detail/:id", controller.GetRoomDetails)
			protected.POST("/room_details", controller.CreateRoomDetail)
			protected.PATCH("/room_details", controller.UpdateRoomDetail)
			protected.DELETE("/room_details/:id", controller.DeleteRoomDetail)

			// RoomDataList Routes
			protected.GET("/room_data_lists", controller.ListRoomDataList)
			protected.GET("/room_data_list/:id", controller.GetRoomDataList)
			protected.POST("/room_data_lists", controller.CreateRoomDataList)
			protected.PATCH("/room_data_lists", controller.UpdateRoomDataList)
			protected.DELETE("/room_data_lists/:id", controller.DeleteRoomDataList)

			// MedicineandVaccine Routes
			protected.GET("/medicineand_vaccines", controller.ListMedicineandVaccines)
			protected.GET("/medicineandvaccine/:id", controller.GetMedicineandVaccine)
			protected.POST("/medicineand_vaccines", controller.CreateMedicineandVaccine)
			protected.PATCH("/medicineand_vaccines", controller.UpdateMedicineandVaccine)
			protected.DELETE("/medicineand_vaccines/:id", controller.DeleteMedicineandVaccine)

			// DosageForm Routes
			protected.GET("/dosage_forms", controller.ListDosageForms)
			protected.GET("/dosageform/:id", controller.GetDosageForm)
			protected.POST("/dosage_forms", controller.CreateDosageForm)
			protected.PATCH("/dosage_forms", controller.UpdateDosageForm)
			protected.DELETE("/dosage_forms/:id", controller.DeleteDosageForm)

			// Age Routes
			protected.GET("/ages", controller.ListAges)
			protected.GET("/age/:id", controller.GetAge)
			protected.POST("/ages", controller.CreateAge)
			protected.PATCH("/ages", controller.UpdateAge)
			protected.DELETE("/ages/:id", controller.DeleteAge)

			// Category Routes
			protected.GET("/categories", controller.ListCategories)
			protected.GET("/category/:id", controller.GetCategory)
			protected.POST("/categories", controller.CreateCategory)
			protected.PATCH("/categories", controller.UpdateCategory)
			protected.DELETE("/categories/:id", controller.DeleteCategory)

			// Appointment Routes
			protected.GET("/appointments", controller.ListAppointment)
			protected.GET("/appointments/:id", controller.GetAppointment)
			protected.POST("/appointments", controller.CreateAppoint)
			protected.PATCH("/appointments", controller.UpdateAppointment)
			protected.DELETE("/appointments/:id", controller.DeleteAppointment)

			// Germ Routes
			protected.GET("/germ", controller.ListGerm)
			protected.GET("/germ/:id", controller.GetGerm)
			protected.POST("/germ", controller.CreateGerm)
			protected.PATCH("/germ", controller.UpdateGerm)
			protected.DELETE("/germ/:id", controller.DeleteGerm)

			// CatchingType Routes
			protected.GET("/catching_type", controller.ListCatchingType)
			protected.GET("/catching_type/:id", controller.GetCatchingType)
			protected.POST("/catching_type", controller.CreateCatchingType)
			protected.PATCH("/catching_type", controller.UpdateCatchingType)
			protected.DELETE("/catching_type/:id", controller.DeleteCatchingType)

			// RiskGroupType Routes
			protected.GET("/risk_group_type", controller.ListRiskGroupType)
			protected.GET("/risk_group_type/:id", controller.GetRiskGroupType)
			protected.POST("/risk_group_type", controller.CreateRiskGroupType)
			protected.PATCH("/risk_group_type", controller.UpdateRiskGroupType)
			protected.DELETE("/risk_group_type/:id", controller.DeleteRiskGroupType)

			// Contagious Routes
			protected.GET("/contagious", controller.ListContagious)
			protected.GET("/contagious/:id", controller.GetContagious)
			protected.POST("/contagious", controller.CreateContagious)
			protected.PATCH("/contagious", controller.UpdateContagious)
			protected.DELETE("/contagious/:id", controller.DeleteContagious)

			//PatientRoom
			protected.GET("/rooms",controller.ListRooms)
			protected.GET("/rooms/:id",controller.GetRoom)
			protected.POST("/rooms",controller.CreateRoom)
			protected.PATCH("/rooms", controller.UpdateRoom)
			protected.DELETE("/rooms/:id", controller.DeleteRoom)
			//Symptom
			protected.GET("/symptoms",controller.ListSymptoms)
			protected.GET("/symptom/:id",controller.GetSymptom)
			protected.POST("/symptoms",controller.CreateSymptom)
			protected.PATCH("/symptoms", controller.UpdateSymptom)
			protected.DELETE("/symptoms/:id", controller.DeleteSymptom)
			//Screening
			protected.GET("/screenings",controller.ListScreening)
			protected.GET("/screening/:id",controller.GetScreening)
			protected.POST("/screenings",controller.CreateScreening)
			protected.PATCH("/screenings", controller.UpdateScreening)
			protected.DELETE("/screenings/:id", controller.DeleteScreening)
		}
	}
	// User Routes
	r.POST("/officers", controller.CreateOfficer)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
