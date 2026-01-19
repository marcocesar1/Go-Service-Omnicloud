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
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/city"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http/responses"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/persistance/persistance_mock"
)

func sendCreatePeopleRequest(newPeople models.People, handler http.HandlerFunc) *httptest.ResponseRecorder {
	router := chi.NewRouter()
	router.Post("/people", handler)

	body, _ := json.Marshal(newPeople)

	req := httptest.NewRequest(
		http.MethodPost,
		"/people",
		strings.NewReader(
			string(body),
		),
	)

	req.Header.Set("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()

	router.ServeHTTP(respRecorder, req)

	return respRecorder
}

func TestCreatePeopleHandler_Success(t *testing.T) {
	cityService := city.NewRandomCityApiMock()
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	usecase := people.NewPeopleUseCase(peopleRepository, cityService)
	handler := (&PeopleHandlers{}).Create(usecase)

	newPeople := models.People{
		Name:  "Test name",
		Email: "test_email@example.com",
	}

	respRecorder := sendCreatePeopleRequest(newPeople, handler)

	if respRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, respRecorder.Code)
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

	if people.ID.Hex() == "" {
		t.Errorf("id field must not be empty")
	}

	if people.Status != models.StatusOut {
		t.Errorf("status field must be %s, got %s", models.StatusOut, people.Status)
	}
}

func TestCreatePeopleHandler_ValidationErrors(t *testing.T) {
	tests := []struct {
		name       string
		people     *models.People
		statusCode int
	}{
		{
			name: "empty name",
			people: &models.People{
				Name:  "",
				Email: "test@test.com",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "short name",
			people: &models.People{
				Name:  "jo",
				Email: "test@test.com",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "empty email",
			people: &models.People{
				Name:  "John",
				Email: "",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			people: &models.People{
				Name:  "John",
				Email: "test",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "duplicated email",
			people: &models.People{
				Name:  "John",
				Email: "johndoe@example.com",
			},
			statusCode: http.StatusConflict,
		},
	}

	cityService := city.NewRandomCityApiMock()
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			peopleRepository.Reset()

			usecase := people.NewPeopleUseCase(peopleRepository, cityService)
			handler := (&PeopleHandlers{}).Create(usecase)

			respRecorder := sendCreatePeopleRequest(*testItem.people, handler)

			if respRecorder.Code != testItem.statusCode {
				t.Fatalf("expected status %d, got %d", testItem.statusCode, respRecorder.Code)
			}
		})
	}
}
