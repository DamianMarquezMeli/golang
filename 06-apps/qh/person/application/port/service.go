package port

import "github.com/devpablocristo/golang/06-apps/qh/person/domain"

type Service interface {
	GetPerson(domain.Person) error
	GetPersonByID(domain.Person) error
	CreatePerson(domain.Person) error
	ListPersons() map[string]domain.Person
}
