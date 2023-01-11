package main

import (
	"os"
	"sync"

	chiAdapter "github.com/devpablocristo/golang/06-apps/qh/person/infrastructure/driver-adapter/handler/chi"
)

const defaultPort = "8080"

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// mux := mux.NewRouter()
	// repo := madb.NewMapDB()
	// service := appl.NewInventoryService(repo)
	// handler := hand.NewHandler(service, port, mux)

	// log.Println("Server listining on port", port)
	// handler.SetupRoutes()
	// err := handler.StartServer()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	wg.Add(2)
	go chiAdapter.People(&wg)
	go chiAdapter.StartApi(&wg, port)
	//go chiAdapter.StartApi(&wg, port)

	// app := fx.New(
	// 	fx.Provide(
	// 		settings.New,
	// 	),
	// 	fx.Invoke(),
	// )

	// app.Run()

}

// func shortner() {
// 	// 1. Create a Repository (redis or mongo)
// 	// repository, err := mongodb.NewMongoRepository("mongodb://localhost:27017", "go_projects", 5)
// 	repository, err := redis.NewRedisRepository("redis://localhost:6379")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	// 2. Instantiate the Application Service
// 	service := shortener.NewRedirectService(repository)

// 	// 3. Instantiate the UseCase that uses the service
// 	useCase := shortener.NewUseCaseShortener(service)

// 	// Wait Group to wait each Goroutine
// 	wg := sync.WaitGroup{}
// 	wg.Add(3)

// 	// 4.1 Create a GRPC Server :7000
// 	grpcHandler := grpcAdapter.NewGrpcHandler(*useCase)
// 	go func() {
// 		grpcHandler.Run(":7000")
// 		wg.Done()
// 	}()

// 	// 4.2 Create a HTTP handler (CHI) :8000
// 	chiHandler := chiAdapter.NewChiHandler(*useCase)
// 	chiHandler.SetupRoutes()
// 	go func() {
// 		chiHandler.Run(":8000")
// 		wg.Done()
// 	}()

// 	// 4.3. Create another HTTP handler (GIN) :9000
// 	ginHandler := ginAdapter.NewGinHandler(*useCase)
// 	ginHandler.SetupRoutes()
// 	go func() {
// 		ginHandler.Run(":9000")
// 		wg.Done()
// 	}()

// 	wg.Wait()
// }
