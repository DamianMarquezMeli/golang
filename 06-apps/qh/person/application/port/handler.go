package port

import "github.com/devpablocristo/golang/06-apps/qh/person/domain"

type Handler interface {
	Exists(uuid string) bool
	Add(p domain.Person)
	List() map[string]domain.Person
}
