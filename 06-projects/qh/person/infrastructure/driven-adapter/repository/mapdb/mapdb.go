package mapdb

import (
	domain "github.com/devpablocristo/golang/06-apps/qh/person/domain"
)

type MapDB struct {
	inmemDB map[string]domain.Person
}

func NewMapDB() *MapDB {
	m := make(map[string]domain.Person)
	return &MapDB{
		inmemDB: m,
	}
}

func (m *MapDB) SavePerson(p domain.Person) error {
	m.inmemDB[p.UUID] = p
	return nil
}

func (m *MapDB) GetPerson(UUID string) (domain.Person, error) {
	return m.inmemDB[UUID], nil
}

func (m MapDB) ListPersons() ([]domain.Person, error) {
	var results []domain.Person
	for _, person := range m.inmemDB {
		results = append(results, person)
	}
	return results, nil
}

func (m MapDB) DeletePerson(UUID string) error {
	delete(m.inmemDB, UUID)
	return nil
}

func (m MapDB) UpdatePerson(UUID string) error {
	return nil
}
