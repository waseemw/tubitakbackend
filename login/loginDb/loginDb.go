package loginDb

import (
	"tubitakPrototypeGo/database"
)

func LoginForPatient(patientTc string, patientId, patientName *string) error {
	err := database.Db.QueryRow("select patient_id,patient_name from patient_table where patient_tc=$1", patientTc).Scan(&*patientId, &*patientName)
	return err
}
