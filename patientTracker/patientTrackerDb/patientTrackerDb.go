package patientTrackerDb

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"tubitakPrototypeGo/database"
)

func SendLocationInfoDb(patientId, beaconId, currentTime string, distance decimal.Decimal) (*sql.Rows, error) {
	_, err := database.Db.Query("insert into patient_tracker_info_table( patient_id, beacon_id, seen_time, distance) values ($1,$2,$3,$4)", patientId, beaconId, currentTime, distance)
	return nil, err

}
