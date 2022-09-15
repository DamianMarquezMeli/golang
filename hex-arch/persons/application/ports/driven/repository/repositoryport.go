package repositoryport

import (
	domain "github.com/devpablocristo/conceptos-go/hex-arch/persons/domain"
)

type Storage interface {
	Exists(uuid string) bool
	Add(p domain.Person)
	List() map[string]domain.Person
}
