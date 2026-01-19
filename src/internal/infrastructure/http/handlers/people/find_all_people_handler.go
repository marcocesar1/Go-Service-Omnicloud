package people

import (
	"fmt"
	"net/http"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/responses"
)

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
