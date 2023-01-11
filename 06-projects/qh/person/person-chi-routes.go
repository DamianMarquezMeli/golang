package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func SetupChiRoutes() http.Handler {
	chiMux := chi.NewRouter()
	//cr.Use("cors")
	//cr.Use(middleware.Logger)

	chiMux.Route("/api/v1", func(r chi.Router) {
		r.Route("/persons", func(r chi.Router) {
			// r.Get("/list-persons", chihandler.GetPersons)
			// r.Post("/create-person", chihandler.CreatePerson)
			// r.Get("/get/{personID}", chihandler.GetPersonByID)
		})
	})

	return chiMux
}
