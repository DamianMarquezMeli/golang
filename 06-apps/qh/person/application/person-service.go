package application

import (
	"github.com/devpablocristo/golang/06-apps/qh/person/domain"
)

type PersonService struct {
	//storage port.Storage
	//handler port.Handler
}

func NewPersonaApplication( /*s port.Storage, h port.Handler*/ ) *PersonService {
	return &PersonService{
		//storage: s,
		//handler: h,
	}
}

func (ps *PersonService) GetPerson(p domain.Person) error {
	// if ps.storage.Exists(p.UUID) {
	// 	return ErrPersonExists
	// }

	// ps.storage.Add(p)

	return nil
}

func (ps *PersonService) GetPersonByID(p domain.Person) error {
	// if ps.storage.Exists(p.UUID) {
	// 	return ErrPersonExists
	// }

	// ps.storage.Add(p)

	return nil
}

func (ps *PersonService) CreatePerson(p domain.Person) error {
	// if ps.storage.Exists(p.UUID) {
	// 	return ErrPersonExists
	// }

	// ps.storage.Add(p)

	return nil
}

func (ps *PersonService) ListPersons() map[string]domain.Person {
	//return ps.storage.List()

	mp := make(map[string]domain.Person)
	return mp
}
