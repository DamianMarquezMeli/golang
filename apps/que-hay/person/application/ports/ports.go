package port

// import "github.com/devpablocristo/go-concepts/hex-arch/persons/domain"

type PersonService interface {
	CreatePerson(p domain.Person) error
	List() map[string]domain.Person

	// 	RunServer(port string)
	// 	SetupRoutes()
}

type Storage interface {
	Exists(uuid string) bool
	Add(p domain.Person)
	List() map[string]domain.Person
}
