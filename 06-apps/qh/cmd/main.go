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

	wg.Add(2)
	go chiAdapter.People(&wg)
	go chiAdapter.StartApi(&wg, port)

	// app := fx.New(
	// 	fx.Provide(
	// 		settings.New,
	// 	),
	// 	fx.Invoke(),
	// )

	// app.Run()
}
