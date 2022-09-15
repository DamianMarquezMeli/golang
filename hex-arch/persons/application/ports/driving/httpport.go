package httpport

import domain "github.com/devpablocristo/go-concepts/hex-arch/persons/domain"

type PersonService interface {
	CreatePerson(p domain.Person) error
	List() map[string]domain.Person
}
