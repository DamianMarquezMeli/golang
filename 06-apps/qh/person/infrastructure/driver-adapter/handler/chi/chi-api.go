package chiAdapter

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/devpablocristo/golang/06-apps/qh/person/application"
	"github.com/devpablocristo/golang/06-apps/qh/person/domain"
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

	ps := application.NewPersonaApplication()
	handler := NewChiHandler(ps, port)

	handler.SetupChiRoutes()
	handler.RunChiServer()
}
