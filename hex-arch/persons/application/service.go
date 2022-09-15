package application

import (
	"errors"

	repositoryport "github.com/devpablocristo/conceptos-go/hex-arch/persons/application/ports/driven/repository"
	domain "github.com/devpablocristo/conceptos-go/hex-arch/persons/domain"
)

var (
	ErrPersonExists = errors.New("person exits")
)

// aqui uso el puerto, o sea la interface
// de esta forma no importa que tipo se pase por parametro
// siempre y cuando cumpla con el principio de polimorfismo
type PersonService struct {
	storage repositoryport.Storage
}

// PersonaApplication no debe tener resposabilidad desde inmemorydb
// Se inyecta en el contructor un "storage", lo que se pide en el struct
// Aqui tb se implementa el puerto
func NewPersonaApplication(s repositoryport.Storage) *PersonService {
	// De esta forma hay inyección de dependecias
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
