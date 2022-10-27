package ports

type HTTPHandler interface {
	SetupRoutes()
	Run() error
}
