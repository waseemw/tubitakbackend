package adminPanel

import (
	"encoding/json"
	"fmt"
	"github.com/barisesen/tcverify"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"tubitakPrototypeGo/adminPanel/adminPanelDatabase"
	"tubitakPrototypeGo/helpers"
)

func SetupAdminPanel(rg *gin.RouterGroup) {
	rg.GET("/all_beacons_admin/:page", getAllBeaconInfo)
	rg.GET("/all_patients_admin/:page", getAllPatientInfo)
	rg.POST("/add_admin", signup)
	rg.POST("/login_admin", login)
	rg.GET("/get_info_patient/:patientId/:page", getSinglePatientTrackingInfo)
	rg.GET("/get_single_patient/:singlePatientId", getSinglePatientInfo)
	rg.GET("/get_info_beacon/:beaconId/:page", getSingleBeaconTrackingInfo)
	rg.POST("/add_relative", addRelative)

}

const itemsPerPage = 10

func login(c *gin.Context) {
	body, err := loginStructFunc(c)
	var password string
	err = adminPanelDatabase.LoginDb(body.Username, &password)
	if err != nil {
		helpers.MyAbort(c, "Admin could not be found")
		return
	}
	passwordTrue := helpers.Checkpassword(body.Password, password)

	if passwordTrue {
		c.JSON(200, "Welcome admin "+body.Username)
		return
	} else {
		helpers.MyAbort(c, "Password or username is wrong")
		return
	}
}

func signup(c *gin.Context) {
	body, err := loginStructFunc(c)
	password, _ := helpers.Hashpassword(body.Password)
	var username string
	err = adminPanelDatabase.SignUpDb(body.Username, password, &username)
	if err != nil {
		//fmt.Println(err)
		helpers.MyAbort(c, "Admin Is already exist")
		return
	}
	c.JSON(200, "Admin "+username+" is added ")

}

func getAllBeaconInfo(c *gin.Context) {
	offSet, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		helpers.MyAbort(c, "Page Number Format is wrong")
		return
	}
	allBeaconsInfoRows, totalBeconNum, err := getAllBeaconRows(offSet * itemsPerPage)
	if err != nil {
		//fmt.Println(err)
		helpers.MyAbort(c, "Could not reach beacons info")
		return
	}
	c.JSON(200, gin.H{"allBeaconsInfoRows": allBeaconsInfoRows,
		"totalBeaconNum": totalBeconNum})
}

func getAllPatientInfo(c *gin.Context) {
	offSet, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		helpers.MyAbort(c, "Page Number Format is wrong")
		return
	}
	allPatientsInfoRows, totalPatientNum, err := getAllPatientRows(offSet * itemsPerPage)
	if err != nil {
		helpers.MyAbort(c, "Could not reach patients info")
		return
	}
	c.JSON(200, gin.H{
		"allPatientsInfoRows": allPatientsInfoRows,
		"totalPatientNum":     totalPatientNum,
	})
}

func getSinglePatientTrackingInfo(c *gin.Context) {
	offSet, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		helpers.MyAbort(c, "Page Number Format is wrong")
		return
	}
	patientId := c.Param("patientId")
	allPatientTrackInfo, totalSinglePatientNum, err := getSinglePatientRows(patientId, offSet*itemsPerPage)
	if err != nil {
		helpers.MyAbort(c, "Something went wrong for "+patientId)
		return
	}
	c.JSON(200, gin.H{
		"totalSinglePatientNum": totalSinglePatientNum,
		"allPatientTrackInfo":   allPatientTrackInfo,
	})

}

func getSingleBeaconTrackingInfo(c *gin.Context) {
	offSet, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		helpers.MyAbort(c, "Page Number Format is wrong")
		return
	}
	beaconId := c.Param("beaconId")
	allBeaconTrackingInfo, singleBeaconTrackNum, err := getSingleBeaconId(beaconId, offSet*itemsPerPage)
	if err != nil {
		helpers.MyAbort(c, "Something went wrong for "+beaconId)
		return
	}
	c.JSON(200, gin.H{
		"totalSingleBeaconTrackNum": singleBeaconTrackNum,
		"allBeaconTrackingInfo":     allBeaconTrackingInfo,
	})

}

func getSinglePatientInfo(c *gin.Context) {
	patientId := c.Param("singlePatientId")
	fmt.Println("patient", patientId)
	row, err := getSinglePatientInfoRow(patientId)
	if err != nil {
		helpers.MyAbort(c, "Patient couldn't be got it")
		return
	}

	c.JSON(200, row)
}

func addRelative(c *gin.Context) {
	body := emailStruct{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Input format is wrong")
		return
	}
	err = json.Unmarshal(data, &body)
	if !helpers.EmailIsValid(body.Email) {
		helpers.MyAbort(c, "Check your email type!!!")
		return
	}
	//check if tc is valid or not
	resp, err := tcverify.Validate(body.PatientTc)
	if err != nil || !resp {
		helpers.MyAbort(c, "Tc is not valid!")
		return
	}

	currentTime := time.Now().Format("2006-01-02 3:4:5 PM")
	password, _ := helpers.GenerateOTP(6)
	//email is got by user
	helpers.SendEmail(password, body.Email)
	token := helpers.TokenGenerator()

	hassPassword, err := helpers.Hashpassword(password)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = adminPanelDatabase.EmailSave(body.Email, hassPassword, currentTime, token, body.PatientTc)
	if err != nil {
		helpers.MyAbort(c, "Relative  is already exist.")
		return
	}

	//sent the current time as well to save on local storage or phone storage.
	//You can save the time on ur database to check it
	c.JSON(200, "Code is successfully sent to "+body.Email)
}
