package main

import (
	github.comgin-gonicgin
	github.comsut64team14controller
	github.comsut64team14entity
	github.comsut64team14middlewares
)

func main() {
	entity.SetupDatabase()

	r = gin.Default()
	r.Use(CORSMiddleware())
	api = r.Group()
	{
		protected = api.Use(middlewares.Authorizes())
		{
			Officer Routes
			protected.GET(officers, controller.ListOfficers)
			protected.GET(officerid, controller.GetOfficer)

			protected.PATCH(officers, controller.UpdateOfficer)
			protected.DELETE(officersid, controller.DeleteOfficer)

			 Patient Routes
			protected.GET(patients, controller.ListPatients)
			protected.GET(patientid, controller.GetPatient)
			protected.POST(patients, controller.CreatePatient)
			protected.PATCH(patients, controller.UpdatePatient)
			protected.DELETE(patientsid, controller.DeletePatient)

			 Specialist Routes
			protected.GET(specialists, controller.ListSpecialists)
			protected.GET(specialistid, controller.GetSpecialist)
			protected.POST(specialists, controller.CreateSpecialist)
			protected.PATCH(specialists, controller.UpdateSpecialist)
			protected.DELETE(specialistsid, controller.DeleteSpecialist)

			 RoomDetail Routes
			protected.GET(room_details, controller.ListRoomDetails)
			protected.GET(room_detailid, controller.GetRoomDetails)
			protected.POST(room_details, controller.CreateRoomDetail)
			protected.PATCH(room_details, controller.UpdateRoomDetail)
			protected.DELETE(room_detailsid, controller.DeleteRoomDetail)

			 RoomDataList Routes
			protected.GET(room_data_lists, controller.ListRoomDataList)
			protected.GET(room_data_listid, controller.GetRoomDataList)
			protected.POST(room_data_lists, controller.CreateRoomDataList)
			protected.PATCH(room_data_lists, controller.UpdateRoomDataList)
			protected.DELETE(room_data_listsid, controller.DeleteRoomDataList)

			 MedicineandVaccine Routes
			protected.GET(medicineand_vaccines, controller.ListMedicineandVaccines)
			protected.GET(medicineandvaccineid, controller.GetMedicineandVaccine)
			protected.POST(medicineand_vaccines, controller.CreateMedicineandVaccine)
			protected.PATCH(medicineand_vaccines, controller.UpdateMedicineandVaccine)
			protected.DELETE(medicineand_vaccinesid, controller.DeleteMedicineandVaccine)

			 DosageForm Routes
			protected.GET(dosage_forms, controller.ListDosageForms)
			protected.GET(dosageformid, controller.GetDosageForm)
			protected.POST(dosage_forms, controller.CreateDosageForm)
			protected.PATCH(dosage_forms, controller.UpdateDosageForm)
			protected.DELETE(dosage_formsid, controller.DeleteDosageForm)

			 Age Routes
			protected.GET(ages, controller.ListAges)
			protected.GET(ageid, controller.GetAge)
			protected.POST(ages, controller.CreateAge)
			protected.PATCH(ages, controller.UpdateAge)
			protected.DELETE(agesid, controller.DeleteAge)

			 Category Routes
			protected.GET(categories, controller.ListCategories)
			protected.GET(categoryid, controller.GetCategory)
			protected.POST(categories, controller.CreateCategory)
			protected.PATCH(categories, controller.UpdateCategory)
			protected.DELETE(categoriesid, controller.DeleteCategory)

			 Appointment Routes
			protected.GET(appointments, controller.ListAppointment)
			protected.GET(appointmentsid, controller.GetAppointment)
			protected.POST(appointments, controller.CreateAppoint)
			protected.PATCH(appointments, controller.UpdateAppointment)
			protected.DELETE(appointmentsid, controller.DeleteAppointment)

			 Germ Routes
			protected.GET(germ, controller.ListGerm)
			protected.GET(germid, controller.GetGerm)
			protected.POST(germ, controller.CreateGerm)
			protected.PATCH(germ, controller.UpdateGerm)
			protected.DELETE(germid, controller.DeleteGerm)

			 CatchingType Routes
			protected.GET(catching_type, controller.ListCatchingType)
			protected.GET(catching_typeid, controller.GetCatchingType)
			protected.POST(catching_type, controller.CreateCatchingType)
			protected.PATCH(catching_type, controller.UpdateCatchingType)
			protected.DELETE(catching_typeid, controller.DeleteCatchingType)

			 RiskGroupType Routes
			protected.GET(risk_group_type, controller.ListRiskGroupType)
			protected.GET(risk_group_typeid, controller.GetRiskGroupType)
			protected.POST(risk_group_type, controller.CreateRiskGroupType)
			protected.PATCH(risk_group_type, controller.UpdateRiskGroupType)
			protected.DELETE(risk_group_typeid, controller.DeleteRiskGroupType)

			 Contagious Routes
			protected.GET(contagious, controller.ListContagious)
			protected.GET(contagiousid, controller.GetContagious)
			protected.POST(contagious, controller.CreateContagious)
			protected.PATCH(contagious, controller.UpdateContagious)
			protected.DELETE(contagiousid, controller.DeleteContagious)

			PatientRoom
			protected.GET(rooms,controller.ListRooms)
			protected.GET(roomsid,controller.GetRoom)
			protected.POST(rooms,controller.CreateRoom)
			protected.PATCH(rooms, controller.UpdateRoom)
			protected.DELETE(roomsid, controller.DeleteRoom)
			Symptom
			protected.GET(symptoms,controller.ListSymptoms)
			protected.GET(symptomid,controller.GetSymptom)
			protected.POST(symptoms,controller.CreateSymptom)
			protected.PATCH(symptoms, controller.UpdateSymptom)
			protected.DELETE(symptomsid, controller.DeleteSymptom)
			Screening
			protected.GET(screenings,controller.ListScreening)
			protected.GET(screeningid,controller.GetScreening)
			protected.POST(screenings,controller.CreateScreening)
			protected.PATCH(screenings, controller.UpdateScreening)
			protected.DELETE(screeningsid, controller.DeleteScreening)

			Prevention
			protected.GET("/preventions", controller.ListPrevention)
			protected.GET("/prevention/:id", controller.GetPrevention)
			protected.POST("/preventions", controller.CreatePrevention)
			protected.PATCH("/preventions", controller.UpdatePrevention)
			protected.DELETE("/preventions/:id", controller.DeletePrevention)
		}
	}
	 User Routes
	r.POST(officers, controller.CreateOfficer)

	 Authentication Routes
	r.POST(login, controller.Login)

	 Run the server
	r.Run()

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c gin.Context) {
		c.Writer.Header().Set(Access-Control-Allow-Origin, )
		c.Writer.Header().Set(Access-Control-Allow-Credentials, true)
		c.Writer.Header().Set(Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With)
		c.Writer.Header().Set(Access-Control-Allow-Methods, POST, OPTIONS, GET, PUT)

		if c.Request.Method == OPTIONS {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}