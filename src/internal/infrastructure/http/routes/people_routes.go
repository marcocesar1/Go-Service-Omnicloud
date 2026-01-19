package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/handlers"
)

type PeopleRoutesInput struct {
	CreatePeopleUseCase *people.CreatePeopleUseCase
	GetPeopleUseCase    *people.GetPeopleUseCase
	GetOnePeopleUseCase *people.GetOnePeopleUseCase
}

type PeopleRoutes struct {
	createPeopleUseCase *people.CreatePeopleUseCase
	getPeopleUseCase    *people.GetPeopleUseCase
	getOnePeopleUseCase *people.GetOnePeopleUseCase
}

func NewPeopleRoutes(input *PeopleRoutesInput) *PeopleRoutes {
	return &PeopleRoutes{
		createPeopleUseCase: input.CreatePeopleUseCase,
		getPeopleUseCase:    input.GetPeopleUseCase,
		getOnePeopleUseCase: input.GetOnePeopleUseCase,
	}
}

func (p *PeopleRoutes) LoadRoutes() *chi.Mux {
	handlers := handlers.CreatePeopleHandlers()

	router := chi.NewRouter()

	router.Get("/", handlers.FindAll(p.getPeopleUseCase))
	router.Post("/", handlers.Create(p.createPeopleUseCase))
	router.Get("/{id}", handlers.FindOne(p.getOnePeopleUseCase))
	router.Patch("/{id}", handlers.Patch())

	return router
}
