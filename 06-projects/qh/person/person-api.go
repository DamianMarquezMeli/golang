package api

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/devpablocristo/golang/06-apps/qh/person/application"
	"github.com/devpablocristo/golang/06-apps/qh/person/domain"
	"github.com/devpablocristo/golang/06-projects/qh/person/infrastructure/driver-adapter/handler/chihandler"
)

var (
	p1 = domain.Person{}
	p2 = domain.Person{}
)

func init() {
	p1 = domain.Person{
		UUID:      "1",
		Firstname: "Homero",
		Lastname:  "Simpson",
		Age:       39,
	}

	p2 = domain.Person{
		UUID:      "2",
		Firstname: "Marge",
		Lastname:  "Simpson",
		Age:       32,
	}
}

func People(wg *sync.WaitGroup) {
	people := []domain.Person{
		p1,
		p2,
	}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

func StartApi(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	// db := postgres.ConnectToDB()
	// defer db.Close()

	ps := application.NewPersonService()
	chihandler.NewChiHandler(ps)

	routes := SetupChiRoutes()
	RunHttpServer(port, routes)
}
