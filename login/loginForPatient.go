package login

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"tubitakPrototypeGo/helpers"
	"tubitakPrototypeGo/login/loginDb"
)

func SetupLogin(rg *gin.RouterGroup) {
	rg.POST("/login", login)

}

func login(c *gin.Context) {
	body := loginStruct{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Input format is wrong")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad Input")
		return
	}
	var patientId, patientName string
	err = loginDb.LoginForPatient(body.PatientTc, &patientId, &patientName)
	if err != nil {
		helpers.MyAbort(c, "There is no such a patient")
		return
	}
	c.JSON(200, gin.H{
		"patientId":   patientId,
		"patientName": patientName,
	})

}
