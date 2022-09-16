package inmemorydb

import (
	"fmt"

	domain "github.com/devpablocristo/go-concepts/hex-arch/persons/domain"
)

///////////////////////////////////////////////////////////////////
// Dependecy Injection
type InmemoryDB struct {
	Storage map[string]domain.Person
}

// Al pasar por parametro el mapa, se cumple con la inyección de dependecias.
// inmemorydb de esta forma NO ES RESPONSABLE de inicializar el mapa.
func NewInmemoryDB(s map[string]domain.Person) *InmemoryDB {
	return &InmemoryDB{
		Storage: s,
	}
}

/////////////////////////////////////////////////////////////////

func (s *InmemoryDB) Exists(uuid string) bool {
	_, found := s.Storage[uuid]
	return found
}

func (s *InmemoryDB) Add(p domain.Person) {

	fmt.Println(p)
	s.Storage[p.UUID] = p
}

func (s *InmemoryDB) List() map[string]domain.Person {
	return s.Storage
}
