package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type DefaultRoutes struct {
}

func NewDefaultRoutes() *DefaultRoutes {
	return &DefaultRoutes{}
}

func (d *DefaultRoutes) LoadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Welcome to the People API",
		})
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	return router
}
