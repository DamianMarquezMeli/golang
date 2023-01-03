package chiAdapter

import "github.com/go-chi/chi"

func (h ChiHandler) SetupChiRoutes() {
	h.chiRouter.Route("/api/v1", func(r chi.Router) {
		r.Route("/persons", func(r chi.Router) {
			r.Get("/list-persons", h.GetPersons)
			r.Post("/create-person", h.CreatePerson)
			r.Get("/get/{personID}", h.GetPersonByID)
		})
	})
}
