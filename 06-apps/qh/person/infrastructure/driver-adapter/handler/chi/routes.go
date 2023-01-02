package chi

import (
	"github.com/go-chi/chi"
)

func PersonChiRoutes(pch *PersonChiHandler) *chi.Mux {
	mux := chi.NewMux()

	//mux.Use(cors)

	mux.Route("/api/v1", func(r1 chi.Router) {
		r1.Route("/persons", func(r2 chi.Router) {
			r2.Get("/get", pch.GetPerson)
			r2.Post("/create", pch.CreatePerson)
			r2.Get("/get/{personID}", pch.GetPersonByID)
		})
	})

	return mux
}
