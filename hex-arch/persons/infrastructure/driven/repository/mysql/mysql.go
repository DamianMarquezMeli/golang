package mysql

import (
	"database/sql"

	"github.com/devpablocristo/go-concepts/hex-arch/persons/domain"
)

type PersonMySQL struct {
	conn *sql.DB
}

func NewPersonMySQL(conn *sql.DB) *PersonMySQL {
	return &PersonMySQL{conn}
}

func (m *PersonMySQL) Exists(uuid string) bool {
	return true
}

func (m *PersonMySQL) Add(p domain.Person) {
}
