package people

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/responses"
)

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
