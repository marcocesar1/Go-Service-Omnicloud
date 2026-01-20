package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/container"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/routes"
)

type Server struct {
	router       *chi.Mux
	appContainer *container.AppContainer
	port         string
}

func NewServer(app *container.AppContainer, port string) *Server {
	return &Server{
		appContainer: app,
		router:       chi.NewRouter(),
		port:         port,
	}
}

func (s *Server) Start() {

	peopleRoutes := routes.NewPeopleRoutes(&routes.PeopleRoutesInput{
		CreatePeopleUseCase:       s.appContainer.CreatePeopleUseCase,
		GetPeopleUseCase:          s.appContainer.GetPeopleUseCase,
		GetOnePeopleUseCase:       s.appContainer.GetOnePeopleUseCase,
		UpdateStatusPeopleUseCase: s.appContainer.UpdateStatusPeopleUseCase,
	})

	defaultRoutes := routes.NewDefaultRoutes()

	s.router.Use(middleware.Logger)

	s.router.Mount("/", defaultRoutes.LoadRoutes())
	s.router.Mount("/people", peopleRoutes.LoadRoutes())

	fmt.Println("Server running on: http://localhost:" + s.port)

	http.ListenAndServe(":3000", s.router)
}
