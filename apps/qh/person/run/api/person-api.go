package api

import (
	"encoding/json"
	"fmt"

	"github.com/devpablocristo/golang/apps/qh/person/domain"
)

var (
	p1 = domain.Person{}
	p2 = domain.Person{}
)

func init() {
	p1 = domain.Person{
		UUID:     "1",
		Name:     "Homero",
		Lastname: "Simpson",
		Age:      39,
	}

	p2 = domain.Person{
		UUID:     "2",
		Name:     "Marge",
		Lastname: "Simpson",
		Age:      32,
	}
}

func people() {
	people := []domain.Person{
		p1,
		p2,
	}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	//mux := chi.NewRouter
	//mux.Use(cors)
}
