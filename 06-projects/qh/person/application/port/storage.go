package port

import "github.com/devpablocristo/golang/06-apps/qh/person/domain"

type Storage interface {
	SavePerson(domain.Person) error
	GetPerson(string) (domain.Person, error)
	ListPersons() ([]domain.Person, error)
	DeletePerson(string) error
	UpdatePerson(string) error
}
