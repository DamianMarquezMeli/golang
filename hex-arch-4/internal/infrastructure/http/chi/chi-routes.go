package chihandler

import (
	"github.com/go-chi/chi/v5"

	service "github.com/devpablocristo/golang/hex-arch-4/internal/core/application"
)

func routes() {
	r := chi.NewMux()

	r.Get("/patients", service.CreatePatient())
	r.Post("/patients", service.CreatePatient())
}
