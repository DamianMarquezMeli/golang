package chiAdapter

import "github.com/go-chi/chi"

func (h ChiHandler) SetupChiRoutes() {
	h.chiRouter.Route("/api/v1", func(r1 chi.Router) {
		r1.Route("/persons", func(r2 chi.Router) {
			r2.Get("/get", h.GetPerson)
			r2.Post("/create", h.CreatePerson)
			r2.Get("/get/{personID}", h.GetPersonByID)
		})
	})
}
