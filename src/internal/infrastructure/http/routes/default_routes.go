package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type DefaultRoutes struct {
}

func CreateDefaultRoutes() *DefaultRoutes {
	return &DefaultRoutes{}
}

func (d *DefaultRoutes) LoadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("App!"))
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	return router
}
