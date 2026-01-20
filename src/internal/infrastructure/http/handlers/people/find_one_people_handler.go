package people

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/responses"
)

func (p *PeopleHandlers) FindOne(usecase *people.GetOnePeopleUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		people, err := usecase.Execute(id)
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

			message := fmt.Sprintf("error finding people: %s", err.Error())
			responses.ErrorResponse(w, message, http.StatusInternalServerError)
			return
		}

		responses.SuccessResponse(w, people, http.StatusOK)
	}
}
