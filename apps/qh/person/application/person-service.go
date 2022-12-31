package application

import (
	"github.com/devpablocristo/golang/apps/qh/person/application/ports"
	"github.com/devpablocristo/golang/apps/qh/person/domain"
)

type PersonService struct {
	storage ports.Storage
}

func NewPersonaApplication(s ports.Storage) *PersonService {
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

func (ps *PersonService) List() map[string]domain.Person {
	return ps.storage.List()
}
