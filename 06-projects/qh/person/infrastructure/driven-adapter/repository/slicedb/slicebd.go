package slicedb

import (
	domain "github.com/devpablocristo/golang/06-apps/qh/person/domain"
)

type SliceDB struct {
	InMemDB []domain.Person
}

func NewSliceDB() *SliceDB {
	return &SliceDB{}
}

func (s SliceDB) SavePerson(person domain.Person) error {
	s.InMemDB = append(s.InMemDB, person)
	return nil
}

func (s SliceDB) GetPerson(UUID string) (domain.Person, error) {
	for _, person := range s.InMemDB {
		if person.UUID == UUID {
			return person, nil
		}
	}
	return domain.Person{}, nil
}

func (s SliceDB) ListPersons() ([]domain.Person, error) {
	return s.InMemDB, nil
}

func (s SliceDB) DeletePerson(UUID string) error {
	for i, person := range s.InMemDB {
		if person.UUID == UUID {
			s.InMemDB = append(s.InMemDB[:i], s.InMemDB[i+1:]...)
			return nil
		}
	}
	return nil
}

func (s SliceDB) UpdatePerson(UUID string) error {
	return nil
}
