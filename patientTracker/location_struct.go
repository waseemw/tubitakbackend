package patientTracker

type locationInfo struct {
	PatientId string
	BeaconId  string
	Distance  string
}

type allBeaconsInfo struct {
	BeaconId         string
	LocationOfBeacon string
}
