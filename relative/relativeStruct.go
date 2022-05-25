package relative

type signRelative struct {
	Email    string
	Password string
}

type changePasswordSt struct {
	OldPassword string
	NewPassword string
}

type addPatientSt struct {
	PatientBd      string
	PRName         string
	PRNum          string
	PRName2        string
	PRNum2         string
	PatientGender  string
	PatientAddress string
	PatientTc      string
	PatientName    string
	PatientSurname string
	PRSurname      string
	PRSurname2     string
}

type singlePatient struct {
	Location      string
	Distance      string
	SeenTime      string
	GoogleMapLink string
}

type getLastLocationSt struct {
	PatientTc      string
	PatientName    string
	PatientSurname string
	LastSeenTime   *string
	Location       *string
	Distance       *string
}
