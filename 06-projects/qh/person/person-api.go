package api

import (
	"encoding/json"
	"fmt"
	"sync"

	application "github.com/devpablocristo/golang/06-projects/qh/person/application"
	domain "github.com/devpablocristo/golang/06-projects/qh/person/domain"
	slicedb "github.com/devpablocristo/golang/06-projects/qh/person/infrastructure/driven-adapter/repository/slicedb"
	chihandler "github.com/devpablocristo/golang/06-projects/qh/person/infrastructure/driver-adapter/handler/chi"
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

func LoadPerons(wg *sync.WaitGroup) {
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

func StartPersonApi(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	// db := postgres.ConnectToDB()
	// defer db.Close()

	//mdb := mapdb.NewMapDB()
	sdb := slicedb.NewSliceDB()
	//pse := application.NewPersonService(mdb)
	pse := application.NewPersonService(sdb)
	han := chihandler.NewChiHandler(pse)
	rou := SetupChiRoutes(han)
	RunHttpServer(port, rou)
}
