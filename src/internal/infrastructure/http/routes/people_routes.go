package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	handlers "github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/handlers/people"
)

type PeopleRoutesInput struct {
	CreatePeopleUseCase       *people.CreatePeopleUseCase
	GetPeopleUseCase          *people.GetPeopleUseCase
	GetOnePeopleUseCase       *people.GetOnePeopleUseCase
	UpdateStatusPeopleUseCase *people.UpdateStatusPeopleUseCase
}

type PeopleRoutes struct {
	createPeopleUseCase       *people.CreatePeopleUseCase
	getPeopleUseCase          *people.GetPeopleUseCase
	getOnePeopleUseCase       *people.GetOnePeopleUseCase
	updateStatusPeopleUseCase *people.UpdateStatusPeopleUseCase
}

func NewPeopleRoutes(input *PeopleRoutesInput) *PeopleRoutes {
	return &PeopleRoutes{
		createPeopleUseCase:       input.CreatePeopleUseCase,
		getPeopleUseCase:          input.GetPeopleUseCase,
		getOnePeopleUseCase:       input.GetOnePeopleUseCase,
		updateStatusPeopleUseCase: input.UpdateStatusPeopleUseCase,
	}
}

func (p *PeopleRoutes) LoadRoutes() *chi.Mux {
	handlers := handlers.CreatePeopleHandlers()

	router := chi.NewRouter()

	router.Get("/", handlers.FindAll(p.getPeopleUseCase))
	router.Post("/", handlers.Create(p.createPeopleUseCase))
	router.Get("/{id}", handlers.FindOne(p.getOnePeopleUseCase))
	router.Patch("/{id}/status", handlers.PatchStatus(p.updateStatusPeopleUseCase))

	return router
}
