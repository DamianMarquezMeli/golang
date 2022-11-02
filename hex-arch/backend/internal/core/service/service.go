package patientservice

import (
	"errors"

	domain "github.com/devpablocristo/golang/hex-arch/internal/core/domain"
	ports "github.com/devpablocristo/golang/hex-arch/internal/core/service/ports"
)

type PatientService struct {
	patientRepository ports.Repository
}

// func New(patientRepository ports.PatientRepository, uidGen uidgen.UIDGen) *service {
func NewPatientService(patientRepository ports.Repository) *PatientService {
	return &PatientService{
		patientRepository: patientRepository,
	}
}

func (ps *PatientService) GetPatient(id string) (domain.Patient, error) {
	patient, err := ps.patientRepository.GetPatient(id)
	if err != nil {
		return domain.Patient{}, errors.New("get patient from repository has failed")
	}

	return patient, nil
}

func (ps *PatientService) CreatePatient(patientPerson domain.Person, doctor domain.Doctor, hospital string, dignosis string) (domain.Patient, error) {
	patient := domain.Patient{
		Patient:   patientPerson,
		Doctor:    doctor,
		Hospital:  hospital,
		Diagnosis: dignosis,
	}

	err := ps.patientRepository.SavePatient(patient)
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}
