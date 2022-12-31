package application

import (
	"github.com/devpablocristo/golang/apps/qh/person/application/port"
	"github.com/devpablocristo/golang/apps/qh/person/domain"
)

type PersonService struct {
	storage port.Storage
}

func NewPersonaApplication(s port.Storage) *PersonService {
	return &PersonService{
		storage: s,
	}
}

func (ps *PersonService) CreatePerson(p domain.Person) error {
	if ps.storage.Exists(p.UUID) {
		return ErrPersonExists
	}

	ps.storage.Add(p)

	return nil
}

func (ps *PersonService) ListPersons() map[string]domain.Person {
	return ps.storage.List()
}
