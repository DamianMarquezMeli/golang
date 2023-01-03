package port

import (
	"context"

	"github.com/devpablocristo/golang/06-apps/qh/person/domain"
)

type Service interface {
	GetPersons(context.Context) ([]domain.Person, error)
	GetPersonByID(domain.Person) error
	CreatePerson(context.Context, *domain.Person) (*domain.Person, error)
	ListPersons() map[string]domain.Person
}
