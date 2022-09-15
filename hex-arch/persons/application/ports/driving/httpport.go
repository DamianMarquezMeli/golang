package httpport

import domain "github.com/devpablocristo/conceptos-go/hex-arch/persons/domain"

type PersonService interface {
	CreatePerson(p domain.Person) error
	List() map[string]domain.Person
}
