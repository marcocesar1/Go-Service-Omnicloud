package people

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/responses"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/persistance/persistance_mock"
)

func sendGetOnePeopleRequest(id string, handler http.HandlerFunc) *httptest.ResponseRecorder {
	router := chi.NewRouter()
	router.Get("/people/{id}", handler)

	req := httptest.NewRequest(
		http.MethodGet,
		"/people/"+id,
		strings.NewReader(
			string(id),
		),
	)

	req.Header.Set("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()

	router.ServeHTTP(respRecorder, req)

	return respRecorder
}

func TestGetOnePeopleHandler_Success(t *testing.T) {
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	usecase := people.NewGetOnePeopleUseCase(peopleRepository)
	handler := (&PeopleHandlers{}).FindOne(usecase)

	respRecorder := sendGetOnePeopleRequest(persistance_mock.TEST_ID1, handler)

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

	var people models.People
	err = json.Unmarshal(jsonbody, &people)
	if err != nil {
		t.Fatalf("invalid response unmarshal")
	}

	if people.ID.Hex() != persistance_mock.TEST_ID1 {
		t.Errorf("id field must be %s, got %s", persistance_mock.TEST_ID1, people.ID.Hex())
	}
}

func TestGetOnePeopleHandler_ValidationErrors(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		statusCode int
	}{
		{
			name:       "invalid object id",
			id:         "invalid",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "people not found",
			id:         "5f8d9f1e2d862c0008e7b2f9",
			statusCode: http.StatusNotFound,
		},
	}

	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			peopleRepository.Reset()

			usecase := people.NewGetOnePeopleUseCase(peopleRepository)
			handler := (&PeopleHandlers{}).FindOne(usecase)

			respRecorder := sendGetOnePeopleRequest(testItem.id, handler)

			if respRecorder.Code != testItem.statusCode {
				t.Fatalf("expected status %d, got %d", testItem.statusCode, respRecorder.Code)
			}
		})
	}
}
