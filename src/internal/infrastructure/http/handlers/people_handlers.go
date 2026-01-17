package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
)

type PeopleHandlers struct {
}

func CreatePeopleHandlers() *PeopleHandlers {
	return &PeopleHandlers{}
}

func (p *PeopleHandlers) FindOne() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FindOne"))
	}
}

func (p *PeopleHandlers) FindAll() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FindAll"))
	}
}

func (p *PeopleHandlers) Create(usecase *people.CreatePeopleUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var people models.People
		err := json.NewDecoder(r.Body).Decode(&people)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			message := fmt.Sprintf("error decoding request body: %s", err.Error())
			json.NewEncoder(w).Encode(map[string]any{
				"message": message,
			})
			return
		}

		err = usecase.Execute(&people)
		if err != nil {
			if errors.Is(err, domain_err.DuplicatedEmail) {
				w.WriteHeader(http.StatusBadRequest)
				message := fmt.Sprintf("email %s already exists", people.Email)
				json.NewEncoder(w).Encode(map[string]any{
					"message": message,
				})
				return
			}

			message := fmt.Sprintf("error creating person: %s", err.Error())
			json.NewEncoder(w).Encode(map[string]any{
				"message": message,
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(people)
	}
}

func (p *PeopleHandlers) Patch() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Update"))
	}
}
