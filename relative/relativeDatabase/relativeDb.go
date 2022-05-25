package relativeDatabase

import (
	"database/sql"
	"fmt"
	"tubitakPrototypeGo/database"
)

func SinglePatient(patientId string, offSet int) (*sql.Rows, error) {
	rows, err := database.Db.Query("select location,distance,seen_time,bdt.google_map_link from patient_tracker_info_table  left join beacon_devices_table as bdt on patient_tracker_info_table.beacon_id = bdt.device_id left join patient_table as pt on pt.patient_id=patient_tracker_info_table.patient_id where patient_tc=$1  order by seen_time desc LIMIT 10 OFFSET $2", patientId, offSet)
	return rows, err
}

//SignPatient DB

func CheckPasswordDb(email string) (string, string, string, bool) {
	var token, password, patientTc string
	var hasPatient bool
	err := database.Db.QueryRow("select token,password,patient_tc,has_patient from patient_relatives_table where email=$1 ", email).Scan(&token, &password, &patientTc, &hasPatient)
	if err != nil {
		return "", "", "", false
	}

	return token, password, patientTc, hasPatient
}

//changePassword Db

func GetPassword(token string) string {
	var password string
	err := database.Db.QueryRow("select password from patient_relatives_table where token=$1", token).Scan(&password)
	if err != nil {
		return ""
	}
	return password
}

func ChangePassword(token, newPassword string) bool {
	var checkChange bool
	err := database.Db.QueryRow("SELECT  1 from changepassword($1,$2,$3)", token, true, newPassword).Scan(&checkChange)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

//add Patient DB

func AddPatient(PatientBd, PRName, PRNum, PRName2, PRNum2, PatientGender, PatientAddress, PatientTc, PatientName, PatientSurname, PRSurname, PRSurname2, token string) error {
	_, err := database.Db.Query("insert into patient_table(patient_bd, patient_relative_name, patient_relative_phone_number, patient_relative_name2, patient_relative_phone_number2, patient_gender, patient_address, patient_tc, patient_name, patient_surname, patient_relative_surname, patient_relative_surname2) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)",
		PatientBd, PRName, PRNum, PRName2, PRNum2, PatientGender, PatientAddress, PatientTc, PatientName, PatientSurname, PRSurname, PRSurname2)
	if err != nil {
		return err
	}
	database.Db.QueryRow("update patient_relatives_table set has_patient=true where token=$1", token)
	return nil

}

//Common Functions

func EmailExistDB(email string) bool {
	var emailExist bool
	err := database.Db.QueryRow("select exists(select 1 from patient_relatives_table where email=$1)", email).Scan(&emailExist)
	if err != nil {
		return false
	}
	return emailExist
}

func TokenExist(token string) bool {
	var emailExist bool
	err := database.Db.QueryRow("select exists(select 1 from patient_relatives_table where token=$1)", token).Scan(&emailExist)
	if err != nil {
		return false
	}
	return emailExist
}

//getLastLocation Db

func GetLastLocationDb(patientTc string) *sql.Row {
	row := database.Db.QueryRow("select patient_tc, patient_name, patient_surname, ptit.seen_time,bdt.location,ptit.distance from patient_table left join patient_tracker_info_table ptit on patient_table.patient_id = ptit.patient_id left join beacon_devices_table bdt  on bdt.device_id=ptit.beacon_id where patient_tc=$1 order by patient_name,ptit.seen_time desc limit 1", patientTc)
	return row

}
