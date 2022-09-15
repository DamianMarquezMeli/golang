package main

import (
	"sync"

	application "github.com/devpablocristo/go-concepts/hex-arch/persons/application"
	domain "github.com/devpablocristo/go-concepts/hex-arch/persons/domain"
	inmemorydb "github.com/devpablocristo/go-concepts/hex-arch/persons/infrastructure/driven/repository/inmemory"
	goriadapter "github.com/devpablocristo/go-concepts/hex-arch/persons/infrastructure/driving/http/gorilla-mux"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	storage := inmemorydb.NewInmemoryDB(make(map[string]domain.Person))
	//mysql := mysql.NewPersonMySQL(nil)
	personService := application.NewPersonaApplication(storage)
	muxHandler := goriadapter.NewGorillaHandler(personService)

	muxHandler.SetupRoutes()
	go func() {
		muxHandler.RunServer("default")
		wg.Done()
	}()

	wg.Wait()

}
