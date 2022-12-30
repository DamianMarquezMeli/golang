package main

import (
	"go.uber.org/fx"

	"github.com/devpablocristo/golang-examples/inventory/settings"
)

func main() {

	app := fx.New(
		fx.Provide(
			settings.New,
		),
		fx.Invoke(),
	)

	app.Run()

}
