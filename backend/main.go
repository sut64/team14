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
			r.GET("/germ", controller.ListGerm)
			r.GET("/germ/:id", controller.GetGerm)
			r.POST("/germ", controller.CreateGerm)
			r.PATCH("/germ", controller.UpdateGerm)
			r.DELETE("/germ/:id", controller.DeleteGerm)

			// CatchingType Routes
			r.GET("/catching_type", controller.ListCatchingType)
			r.GET("/catching_type/:id", controller.GetCatchingType)
			r.POST("/catching_type", controller.CreateCatchingType)
			r.PATCH("/catching_type", controller.UpdateCatchingType)
			r.DELETE("/catching_type/:id", controller.DeleteCatchingType)

			// RiskGroupType Routes
			r.GET("/risk_group_type", controller.ListRiskGroupType)
			r.GET("/risk_group_type/:id", controller.GetRiskGroupType)
			r.POST("/risk_group_type", controller.CreateRiskGroupType)
			r.PATCH("/risk_group_type", controller.UpdateRiskGroupType)
			r.DELETE("/risk_group_type/:id", controller.DeleteRiskGroupType)

			// Contagious Routes
			r.GET("/contagious", controller.ListContagious)
			r.GET("/contagious/:id", controller.GetContagious)
			r.POST("/contagious", controller.CreateContagious)
			r.PATCH("/contagious", controller.UpdateContagious)
			r.DELETE("/contagious/:id", controller.DeleteContagious)
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
