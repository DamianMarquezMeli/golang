package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"

	appl "github.com/devpablocristo/golang/06-apps/bookstore/inventory/application"
	hand "github.com/devpablocristo/golang/06-apps/bookstore/inventory/infrastructure/handler"
	madb "github.com/devpablocristo/golang/06-apps/bookstore/inventory/infrastructure/repository/inmemory/mapdb"
)

const defaultPort = ":8080"

// var (
// 	mapDB   = mapdb.NewMapDB()
// 	sliceDB = slicedb.NewSliceDB()
// )

// func init() {
// 	book1 := domain.Book{
// 		Author: pdomain.Person{
// 			Firstname: "Isaac",
// 			Lastname:  "Asimov",
// 		},
// 		Title: "Fundation",
// 		Price: 28.50,
// 		ISBN:  "0-553-29335-4",
// 	}

// 	book2 := domain.Book{
// 		Author: pdomain.Person{
// 			Firstname: "Stanislaw",
// 			Lastname:  "Lem",
// 		},
// 		Title: "Solaris",
// 		Price: 65.20,
// 		ISBN:  "0156027607",
// 	}

// 	book3 := domain.Book{
// 		Author: pdomain.Person{
// 			Firstname: "Arthur C.",
// 			Lastname:  "Clarck",
// 		},
// 		Title: "Rendezvous with Rama",
// 		Price: 53.50,
// 		ISBN:  "0-575-01587-X",
// 	}

// 	book4 := domain.Book{
// 		Author: pdomain.Person{
// 			Firstname: "Jorge Luis",
// 			Lastname:  "Borges",
// 		},
// 		Title: "El Aleph",
// 		Price: 42.75,
// 		ISBN:  "84-206-1933-7",
// 	}

// 	domain.Inventory = []domain.BookStock{
// 		{
// 			Book:  book1,
// 			Stock: 41,
// 		},
// 		{
// 			Book:  book2,
// 			Stock: 32,
// 		},
// 		{
// 			Book:  book3,
// 			Stock: 12,
// 		},
// 		{
// 			Book:  book4,
// 			Stock: 93,
// 		},
// 	}

// 	mapDB.SaveBook(&b1)
// 	mapDB.SaveBook(&b2)

// 	sliceDB.SaveBook(b1)
// 	sliceDB.SaveBook(b2)

// 	fmt.Println(*mapDB)
// 	fmt.Println(sliceDB)
// }

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	mux := mux.NewRouter()
	repo := madb.NewMapDB()
	service := appl.NewInventoryService(repo)
	handler := hand.NewHandler(service, port, mux)

	log.Println("Server listining on port", port)
	handler.SetupRoutes()
	err := handler.StartServer()
	if err != nil {
		log.Fatal(err.Error())
	}

}
