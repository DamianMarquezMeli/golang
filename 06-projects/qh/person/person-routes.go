package api

import (
	"net/http"

	"github.com/go-chi/chi"

	chihandler "github.com/devpablocristo/golang/06-projects/qh/person/infrastructure/driver-adapter/handler/chi"
)

func SetupChiRoutes(handler *chihandler.ChiHandler) http.Handler {
	chiMux := chi.NewRouter()
	//chiMux.Use("cors")
	//chiMux.Use(middleware.Logger)

	chiMux.Route("/api/v1", func(r chi.Router) {
		r.Route("/person", func(r chi.Router) {
			r.Post("/create", handler.CreatePerson)
			r.Get("/list", handler.GetPersons)
			r.Get("/get/{personUUID}", handler.GetPerson)
			r.Put("/update", handler.UpdatePerson)
			r.Delete("/delete", handler.DeletePerson)
		})
	})

	return chiMux
}
