package app

import (
	"github.com/devpablocristo/users/controllers/ping"
	"github.com/devpablocristo/users/controllers/users"
)

func urlMap() {
	router.GET("/ping", ping.Ping)

	/*
		ejemplo de dato tipo user, poner en el body y listo
		{
			"user": "ccc1",
			"password": "1234",
			"rol": "agente"
		}
	*/
	router.POST("/users", users.CreateUser)
	router.GET("/users", users.GetUsers)    // Read all
	router.GET("/users/:id", users.GetUser) // Read 1
	router.PUT("/users/:id", users.UpdateUser)
	router.DELETE("/users/:id", users.DeleteUser)
}
