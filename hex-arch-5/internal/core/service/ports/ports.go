package ports

import "github.com/devpablocristo/golang/hex-arch-5/internal/core/domain"

// patient http handler port
type PatientServicePort interface {
	GetPatient(id string) (domain.Patient, error)
	CreatePatient(patient domain.Person, doctor domain.Doctor, hospital string, dignosis string) (domain.Patient, error)
}

// patient repository port
type PatientRepositoryPort interface {
	GetPatient(id string) (domain.Patient, error)
	SavePatient(domain.Patient) error
}

type PatientScrapperPort interface {
	Check() error
}
