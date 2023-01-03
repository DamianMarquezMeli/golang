package handler

import (
	"fmt"

	"github.com/devpablocristo/golang/06-apps/bookstore/inventory/domain"
)

var (
	mapDB   = mapdb.NewMapDB()
	sliceDB = slicedb.NewSliceDB()
)

func init() {
	b1 := domain.Book{
		Author: domain.Person{
			Firstname: "J.K.",
			Lastname:  "Rowling",
		},
		Title: "Harry Potter and the Philosopher's Stone",
		Price: 45.00,
		ISBN:  "hpotter",
	}

	b2 := domain.Book{
		Author: domain.Person{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Foundation",
		Price: 25.24,
		ISBN:  "fasimov",
	}

	mapDB.SaveBook(b1)
	mapDB.SaveBook(b2)

	sliceDB.SaveBook(b1)
	sliceDB.SaveBook(b2)

	fmt.Println(*mapDB)
	fmt.Println(sliceDB)
}
