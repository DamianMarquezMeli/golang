package port

import "github.com/devpablocristo/golang/apps/qh/person/domain"

type PersonService interface {
	CreatePerson(domain.Person) error
	List() map[string]domain.Person
}
