package relative

import (
	"encoding/json"
	"fmt"
	"github.com/barisesen/tcverify"
	"github.com/gin-gonic/gin"
	"strconv"
	"tubitakPrototypeGo/helpers"
	"tubitakPrototypeGo/relative/relativeDatabase"
)

func SetupPatientRelative(rg *gin.RouterGroup) {
	rg.POST("/sign_in", signPatient)
	rg.POST("/change_password", changePassword)
	rg.POST("/add_patient", addPatient)
	rg.GET("/patient_tracking_info/:patientId/:page", getPatientTrackingInfo)
	rg.GET("/last_location/:patientId", getLastLocation)

}

func signPatient(c *gin.Context) {
	body := signRelative{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Admin could not be found")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad Input")
		return
	}
	if !helpers.EmailIsValid(body.Email) {
		helpers.MyAbort(c, "Email formatı yanlış!")
		return
	}

	emailExist := relativeDatabase.EmailExistDB(body.Email)
	if !emailExist {
		helpers.MyAbort(c, "Girmis oldugunuz mail adresi gecerli degildir.")
		return
	}
	token, password, patientTc, hasPatient := relativeDatabase.CheckPasswordDb(body.Email)
	if !helpers.Checkpassword(body.Password, password) {
		helpers.MyAbort(c, "Parola Yanlış!")
		return
	}

	c.JSON(200, gin.H{
		"token":      token,
		"patientTc":  patientTc,
		"hasPatient": hasPatient,
	})

}

func changePassword(c *gin.Context) {
	body := changePasswordSt{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Admin could not be found")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad Input")
		return
	}
	token := c.GetHeader("token")

	password := relativeDatabase.GetPassword(token)
	if !helpers.Checkpassword(body.OldPassword, password) {
		helpers.MyAbort(c, "Eski Parolanin dogru oldugundan emin olun. ")
		return
	} else {
		newPassword, _ := helpers.Hashpassword(body.NewPassword)
		checkChange := relativeDatabase.ChangePassword(token, newPassword)
		if !checkChange {
			helpers.MyAbort(c, "Eski Parolanin dogru oldugundan emin olun. ")
			return
		}
	}

	c.JSON(200, "Parola basariyla degistirilmistir.")

}

func addPatient(c *gin.Context) {
	body := addPatientSt{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Admin could not be found")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad Input")
		return
	}
	token := c.GetHeader("token")
	checkTokenExist := relativeDatabase.TokenExist(token)
	if !checkTokenExist {
		helpers.MyAbort(c, "Birseyler hatali gitti lutfen yeniden baglanin!")
		return
	}
	resp, err := tcverify.Validate(body.PatientTc)
	if err != nil || !resp {
		helpers.MyAbort(c, "Tc is not valid!")
		return
	}

	err = relativeDatabase.AddPatient(body.PatientBd, body.PRName, body.PRNum, body.PRName2, body.PRNum2, body.PatientGender, body.PatientAddress, body.PatientTc, body.PatientName, body.PatientSurname, body.PRSurname, body.PRSurname2, token)
	if err != nil {
		helpers.MyAbort(c, "Hasta onceden kayit edilmistir!")
		return
	}
	c.JSON(200, "Patient Is Added")

}

const itemsPerPage = 10

func getPatientTrackingInfo(c *gin.Context) {
	token := c.GetHeader("token")
	checkTokenExist := relativeDatabase.TokenExist(token)
	if !checkTokenExist {
		helpers.MyAbort(c, "Birseyler hatali gitti lutfen yeniden baglanin!")
		return
	}
	offSet, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		helpers.MyAbort(c, "Page Number Format is wrong")
		return
	}
	patientId := c.Param("patientId")

	allPatientTrackInfo, err := getSinglePatientRows(patientId, offSet*itemsPerPage)
	if err != nil {
		helpers.MyAbort(c, "Something went wrong for "+patientId)
		return
	}
	c.JSON(200, allPatientTrackInfo)

}

func getLastLocation(c *gin.Context) {
	token := c.GetHeader("token")
	checkTokenExist := relativeDatabase.TokenExist(token)
	if !checkTokenExist {
		helpers.MyAbort(c, "Birseyler hatali gitti lutfen yeniden baglanin!")
		return
	}
	patientId := c.Param("patientId")
	lastLocation, err := getLastLocationRow(patientId)
	if err != nil {
		fmt.Println(err)
		helpers.MyAbort(c, "Hastanin kayitli bilgisi bulunamamistir.")
		return
	}
	c.JSON(200, lastLocation)

}
