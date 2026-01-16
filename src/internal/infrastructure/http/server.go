package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/routes"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	return &Server{
		router: chi.NewRouter(),
	}
}

func (s *Server) Start() {

	peopleRoutes := routes.CreatePeopleRoutes()
	defaultRoutes := routes.CreateDefaultRoutes()

	s.router.Use(middleware.Logger)

	s.router.Mount("/", defaultRoutes.LoadRoutes())
	s.router.Mount("/people", peopleRoutes.LoadRoutes())

	fmt.Println("Server running on: http://localhost:3000")

	http.ListenAndServe(":3000", s.router)
}
