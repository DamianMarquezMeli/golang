package chihandler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func routes() http.Handler {
	mux := chi.NewRouter

	mux.Use(cors)
}
