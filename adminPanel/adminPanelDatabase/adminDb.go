package adminPanelDatabase

import (
	"database/sql"
	"tubitakPrototypeGo/database"
)

//adminDb is used for admin database calls.

// LoginDb loginDb
func LoginDb(username string, password *string) error {
	err := database.Db.QueryRow("select admin_password  from admin_table where admin_username=$1", username).Scan(&*password)
	return err
}
func SignUpDb(username, password string, savedUsername *string) error {
	err := database.Db.QueryRow("insert into admin_table(admin_username, admin_password) VALUES($1,$2) returning admin_username", username, password).Scan(&*savedUsername)
	return err

}

//Patient Calls

func GetAllPatientDb(offSet int) (*sql.Rows, int, error) {
	var totalPatientNum int
	rows, err := database.Db.Query("select distinct on (patient_name) patient_tc, patient_name, patient_surname, patient_gender, bdt.location,ptit.seen_time from patient_table left join patient_tracker_info_table ptit on patient_table.patient_id = ptit.patient_id left join beacon_devices_table bdt  on bdt.device_id=ptit.beacon_id order by patient_name,ptit.seen_time desc nulls last, 1 LIMIT 10 OFFSET $1", offSet)
	err = database.Db.QueryRow("select count(patient_tc) from patient_table").Scan(&totalPatientNum)
	return rows, totalPatientNum, err
}

func GetSinglePatientRowsDb(patientId string, offSet int) (*sql.Rows, int, error) {
	var totalSinglePatientNum int
	rows, err := database.Db.Query("select bdt.device_id,location,distance,seen_time,bdt.google_map_link from patient_tracker_info_table  left join beacon_devices_table as bdt on patient_tracker_info_table.beacon_id = bdt.device_id left join patient_table as pt on pt.patient_id=patient_tracker_info_table.patient_id where patient_tc=$1  order by seen_time desc LIMIT 10 OFFSET $2", patientId, offSet)
	err = database.Db.QueryRow("select count(bdt.device_id) from patient_tracker_info_table left join beacon_devices_table as bdt on patient_tracker_info_table.beacon_id = bdt.device_id left join patient_table as pt on pt.patient_id = patient_tracker_info_table.patient_id where patient_tc = '10952336010'").Scan(&totalSinglePatientNum)

	return rows, totalSinglePatientNum, err
}

func GetSinglePatientAllInfo(patientId string) *sql.Row {
	rows := database.Db.QueryRow("select patient_tc,patient_name,patient_surname,patient_bd,patient_relative_name,patient_relative_surname,patient_relative_phone_number,patient_relative_name2,patient_relative_surname2,patient_relative_phone_number2,patient_gender,patient_address from patient_table where patient_tc=$1", patientId)
	return rows
}

//Devices Calls
func GetAllBeaconRowsDb(pageOffset int) (*sql.Rows, int, error) {
	rows, err := database.Db.Query("select * from beacon_devices_table LIMIT 10 OFFSET $1", pageOffset)
	var totalBeaconNum int
	err = database.Db.QueryRow("select count(device_id) from beacon_devices_table").Scan(&totalBeaconNum)

	return rows, totalBeaconNum, err
}

func GetSingleBeaconIdDb(beaconId string, offSet int) (*sql.Rows, int, error) {
	var singleBeaconTrackNum int
	rows, err := database.Db.Query("select pt.patient_tc,seen_time,distance,bdt.location,bdt.google_map_link,bdt.minor,bdt.major from patient_tracker_info_table left join patient_table pt on patient_tracker_info_table.patient_id = pt.patient_id left join beacon_devices_table as bdt  on  bdt.device_id=patient_tracker_info_table.beacon_id where beacon_id=$1 order by seen_time LIMIT 10 OFFSET $2", beaconId, offSet)
	err = database.Db.QueryRow("select count(pt.patient_tc) from patient_tracker_info_table left join patient_table pt on patient_tracker_info_table.patient_id = pt.patient_id left join beacon_devices_table as bdt  on  bdt.device_id=patient_tracker_info_table.beacon_id where beacon_id=$1", beaconId).Scan(&singleBeaconTrackNum)
	return rows, singleBeaconTrackNum, err
}

//email save

func EmailSave(email, password, currentTime, token, patientTc string) error {
	_, err := database.Db.Query("insert into patient_relatives_table(email, password, send_date, token,patient_tc) values ($1,$2,$3,$4,$5)", email, password, currentTime, token, patientTc)
	if err != nil {
		return err
	}
	return nil

}
