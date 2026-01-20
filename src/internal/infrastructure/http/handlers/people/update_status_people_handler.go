package people

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/responses"
)

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
			if errors.Is(err, domain_err.ErrInvalidObjectId) {
				message := fmt.Sprintf("invalid object id: %s", id)
				responses.ErrorResponse(w, message, http.StatusBadRequest)
				return
			}

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
