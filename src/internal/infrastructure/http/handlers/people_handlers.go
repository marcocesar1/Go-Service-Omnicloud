package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
)

type PeopleHandlers struct {
}

func CreatePeopleHandlers() *PeopleHandlers {
	return &PeopleHandlers{}
}

func (p *PeopleHandlers) FindOne(usecase *people.GetOnePeopleUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		w.Header().Set("Content-Type", "application/json")

		people, err := usecase.Execute(id)
		if err != nil {
			if errors.Is(err, domain_err.ErrNotFound) {
				message := fmt.Sprintf("people with id %s not found", id)
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]any{
					"message": message,
				})
				return
			}

			message := fmt.Sprintf("error finding people: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"message": message,
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(people)
	}
}

func (p *PeopleHandlers) FindAll(usecase *people.GetPeopleUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		query := r.URL.Query()
		status := query.Get("status")

		people, err := usecase.Execute(&repositories.FindAllPeopleFilter{
			Status: status,
		})
		if err != nil {
			message := fmt.Sprintf("error finding people: %s", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"message": message,
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(people)
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

func (p *PeopleHandlers) PatchStatus(usecase *people.UpdateStatusPeopleUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := chi.URLParam(r, "id")

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

		peopleUpd, err := usecase.Execute(id, people.Status)
		if err != nil {
			if errors.Is(err, domain_err.ErrNotFound) {
				message := fmt.Sprintf("people with id %s not found", id)
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]any{
					"message": message,
				})
				return
			}

			if errors.Is(err, domain_err.InvalidStatus) {
				message := "invalid status, valid statuses [IN, OUT]"
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]any{
					"message": message,
				})
				return
			}

			if errors.Is(err, domain_err.StatusIsTheSame) {
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(map[string]any{
					"message": err.Error(),
				})
				return
			}

			message := fmt.Sprintf("error updating people: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"message": message,
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"id":         peopleUpd.ID,
			"status":     peopleUpd.Status,
			"updated_at": peopleUpd.UpdatedAt,
		})
	}
}
