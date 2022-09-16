package main

import (
	"sync"

	application "github.com/devpablocristo/go-concepts/hex-arch/persons/application"
	domain "github.com/devpablocristo/go-concepts/hex-arch/persons/domain"
	inmemorydb "github.com/devpablocristo/go-concepts/hex-arch/persons/infrastructure/driven/repository/inmemory"
	ginAdapter "github.com/devpablocristo/go-concepts/hex-arch/persons/infrastructure/driving/http/gin"
	goriAdapter "github.com/devpablocristo/go-concepts/hex-arch/persons/infrastructure/driving/http/gorilla-mux"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	storage := inmemorydb.NewInmemoryDB(make(map[string]domain.Person))
	//mysql := mysql.NewPersonMySQL(nil)
	personService := application.NewPersonaApplication(storage)

	muxHandler := goriAdapter.NewGorillaHandler(personService)
	muxHandler.SetupRoutes()
	go func() {
		muxHandler.RunServer("default")
		wg.Done()
	}()

	ginHandler := ginAdapter.NewGinHandler(personService)
	ginHandler.SetupRoutes()
	go func() {
		ginHandler.Run("default")
		wg.Done()
	}()

	wg.Wait()

}
