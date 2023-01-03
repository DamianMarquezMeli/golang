package handler

import "sync"

var (
	mapDB                             = mapdb.NewMapDB()
	sliceDB                           = slicedb.NewSliceDB()
	httpMuxRouter muxrouter.MuxRouter = *muxrouter.NewMuxRouter()
)

func init() {
	b1 := inventory.Book{
		Author: inventory.Person{
			Firstname: "J.K.",
			Lastname:  "Rowling",
		},
		Title: "Harry Potter and the Philosopher's Stone",
		Price: 45.00,
		ISBN:  "hpotter",
	}

	b2 := inventory.Book{
		Author: inventory.Person{
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

	//fmt.Println(*mapDB)

}
func InventoryService(wg *sync.WaitGroup) {
	defer wg.Done()

	inventoryControllers := getInventoryControllers()

	httpMuxRouter.POST("/inventory/add", inventoryControllers.Add)
	httpMuxRouter.GET("/inventory/all", inventoryControllers.GetAll)
	httpMuxRouter.SERVE(":8888")
}

func getInventoryControllers() http.HTTPInteractor {
	inventoryRepository := repository.NewRepositoryInteractor(mapDB)
	inventoryUseCases := usecases.MakeUseCasesInteractor(inventoryRepository)
	inventoryControllers := http.NewHTTPInteractor(inventoryUseCases)

	return *inventoryControllers
}
