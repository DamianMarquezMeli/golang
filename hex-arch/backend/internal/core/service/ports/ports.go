package ports

import "github.com/devpablocristo/golang/hex-arch/internal/core/domain"

// patient http handler port
type Service interface {
	GetPatient(id string) (domain.Patient, error)
	CreatePatient(patient domain.Person, doctor domain.Doctor, hospital string, dignosis string) (domain.Patient, error)
}

// patient repository port
//
//go:generate mockgen -source=./service.go -destination=../../../mocks/service_mock.go -package=mocks
type Repository interface {
	GetPatient(id string) (domain.Patient, error)
	SavePatient(domain.Patient) error
}

//go:generate mockgen -source=./service.go -destination=../../../mocks/service_mock.go -package=mocks
type Scrapper interface {
	Check() error
}
