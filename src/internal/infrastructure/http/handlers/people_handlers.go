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
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/responses"
)

type PeopleHandlers struct {
}

func CreatePeopleHandlers() *PeopleHandlers {
	return &PeopleHandlers{}
}

func (p *PeopleHandlers) FindOne(usecase *people.GetOnePeopleUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		people, err := usecase.Execute(id)
		if err != nil {
			if errors.Is(err, domain_err.ErrNotFound) {
				message := fmt.Sprintf("people with id %s not found", id)
				responses.ErrorResponse(w, message, http.StatusNotFound)
				return
			}

			message := fmt.Sprintf("error finding people: %s", err.Error())
			responses.ErrorResponse(w, message, http.StatusInternalServerError)
			return
		}

		responses.SuccessResponse(w, people, http.StatusOK)
	}
}

func (p *PeopleHandlers) FindAll(usecase *people.GetPeopleUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		status := query.Get("status")

		people, err := usecase.Execute(&repositories.FindAllPeopleFilter{
			Status: status,
		})
		if err != nil {
			message := fmt.Sprintf("error finding people: %s", err.Error())
			responses.ErrorResponse(w, message, http.StatusInternalServerError)
			return
		}

		responses.SuccessResponse(w, people, http.StatusOK)
	}
}

func (p *PeopleHandlers) Create(usecase *people.CreatePeopleUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var people models.People
		err := json.NewDecoder(r.Body).Decode(&people)
		if err != nil {
			message := fmt.Sprintf("error decoding request body: %s", err.Error())
			responses.ErrorResponse(w, message, http.StatusBadRequest)
			return
		}

		err = usecase.Execute(&people)
		if err != nil {
			if errors.Is(err, domain_err.InvalidPeopleField) {
				message := err.Error()
				responses.ErrorResponse(w, message, http.StatusBadRequest)
				return
			}

			if errors.Is(err, domain_err.DuplicatedEmail) {
				message := fmt.Sprintf("email %s already exists", people.Email)
				responses.ErrorResponse(w, message, http.StatusConflict)
				return
			}

			message := fmt.Sprintf("error creating person: %s", err.Error())
			responses.ErrorResponse(w, message, http.StatusInternalServerError)
			return
		}

		responses.SuccessResponse(w, people, http.StatusCreated)
	}
}

func (p *PeopleHandlers) PatchStatus(usecase *people.UpdateStatusPeopleUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		var people models.People
		err := json.NewDecoder(r.Body).Decode(&people)
		if err != nil {
			message := fmt.Sprintf("error decoding request body: %s", err.Error())
			responses.ErrorResponse(w, message, http.StatusBadRequest)
			return
		}

		peopleUpd, err := usecase.Execute(id, people.Status)
		if err != nil {
			if errors.Is(err, domain_err.ErrNotFound) {
				message := fmt.Sprintf("people with id %s not found", id)
				responses.ErrorResponse(w, message, http.StatusNotFound)
				return
			}

			if errors.Is(err, domain_err.InvalidStatus) {
				message := "invalid status, valid statuses [IN, OUT]"
				responses.ErrorResponse(w, message, http.StatusBadRequest)
				return
			}

			if errors.Is(err, domain_err.StatusIsTheSame) {
				message := err.Error()
				responses.ErrorResponse(w, message, http.StatusConflict)
				return
			}

			message := fmt.Sprintf("error updating people: %s", err.Error())
			responses.ErrorResponse(w, message, http.StatusInternalServerError)
			return
		}

		data := map[string]any{
			"id":         peopleUpd.ID,
			"status":     peopleUpd.Status,
			"updated_at": peopleUpd.UpdatedAt,
		}
		responses.SuccessResponse(w, data, http.StatusOK)
	}
}
