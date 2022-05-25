package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tubitakPrototypeGo/adminPanel"
	"tubitakPrototypeGo/database"
	"tubitakPrototypeGo/login"
	patientTracker2 "tubitakPrototypeGo/patientTracker"
	"tubitakPrototypeGo/relative"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	database.ConnectDatabase() // connection starts at the beginning
	router.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "*")
		context.Header("Access-Control-Allow-Methods", "*")
		if context.Request.Method == "OPTIONS" {
			context.Status(200)
			context.Abort()
		}
	})

	//this is for the patient info
	patient := router.Group("/patient")
	login.SetupLogin(patient)

	//this is for patient tracking. It sends the patient travelling information catching by beacon.
	patientTracker := router.Group("/track")
	patientTracker2.PatientTrackerSetup(patientTracker)

	admin := router.Group("/admin")
	adminPanel.SetupAdminPanel(admin)

	relatives := router.Group("/relative")
	relative.SetupPatientRelative(relatives)

	err := router.Run(":8000")
	if err != nil {
		fmt.Println("Connection can not be completed!")
		return
	}
}
