package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/handlers"
)

type PeopleRoutes struct {
}

func CreatePeopleRoutes() *PeopleRoutes {
	return &PeopleRoutes{}
}

func (p *PeopleRoutes) LoadRoutes() *chi.Mux {
	handlers := handlers.CreatePeopleHandlers()

	router := chi.NewRouter()

	router.Get("/", handlers.FindAll())
	router.Post("/", handlers.Create())
	router.Get("/{id}", handlers.FindOne())
	router.Patch("/{id}", handlers.Patch())

	return router
}
