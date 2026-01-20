package people

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/responses"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/persistance/persistance_mock"
)

func sendGetPeopleRequest(filter *repositories.FindAllPeopleFilter, handler http.HandlerFunc) *httptest.ResponseRecorder {
	router := chi.NewRouter()
	router.Get("/people", handler)

	query := ""
	if filter != nil {
		query = fmt.Sprintf("status=%s", filter.Status)
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/people?"+query,
		nil,
	)

	req.Header.Set("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()

	router.ServeHTTP(respRecorder, req)

	return respRecorder
}

func TestGetPeopleHandler_Success(t *testing.T) {
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	usecase := people.NewGetPeopleUseCase(peopleRepository)
	handler := (&PeopleHandlers{}).FindAll(usecase)

	respRecorder := sendGetPeopleRequest(&repositories.FindAllPeopleFilter{
		Status: string(models.StatusIn),
	}, handler)

	if respRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, respRecorder.Code)
	}

	var response responses.SuccessResponseData
	err := json.NewDecoder(respRecorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("invalid response body")
	}

	jsonbody, err := json.Marshal(response.Data)
	if err != nil {
		t.Fatalf("invalid response marshal")
	}

	var people []models.People
	err = json.Unmarshal(jsonbody, &people)
	if err != nil {
		t.Fatalf("invalid response unmarshal")
	}

	if len(people) != 1 {
		t.Errorf("people length must be 1, got %d", len(people))
	}
}

func TestGetPeopleHandler_Error(t *testing.T) {
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()
	peopleRepository.Error = errors.New("db error mock test")

	usecase := people.NewGetPeopleUseCase(peopleRepository)
	handler := (&PeopleHandlers{}).FindAll(usecase)

	respRecorder := sendGetPeopleRequest(&repositories.FindAllPeopleFilter{
		Status: string(models.StatusIn),
	}, handler)

	if respRecorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusOK, respRecorder.Code)
	}
}
