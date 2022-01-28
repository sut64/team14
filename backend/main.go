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
