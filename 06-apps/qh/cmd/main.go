package main

import (
	"os"
	"sync"

	"github.com/devpablocristo/golang/06-apps/qh/person/infrastructure/driver-adapter/handler/chi"
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
	go chi.People(&wg)
	go chi.StartApi(&wg, port)

	// app := fx.New(
	// 	fx.Provide(
	// 		settings.New,
	// 	),
	// 	fx.Invoke(),
	// )

	// app.Run()
}
