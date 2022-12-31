package main

import (
	"sync"

	"go.uber.org/fx"

	"github.com/devpablocristo/golang/apps/qh/internal/commons/settings"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go wg.Wait()

	app := fx.New(
		fx.Provide(
			settings.New,
		),
		fx.Invoke(),
	)

	app.Run()
}
