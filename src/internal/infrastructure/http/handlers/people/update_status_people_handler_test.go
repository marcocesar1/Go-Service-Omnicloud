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

func sendUpdateStatusPeopleRequest(id string, newPeople models.People, handler http.HandlerFunc) *httptest.ResponseRecorder {
	router := chi.NewRouter()
	router.Patch("/people/{id}", handler)

	body, _ := json.Marshal(newPeople)

	req := httptest.NewRequest(
		http.MethodPatch,
		"/people/"+id,
		strings.NewReader(
			string(body),
		),
	)

	req.Header.Set("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()

	router.ServeHTTP(respRecorder, req)

	return respRecorder
}

func TestUpdateStatusPeopleHandler_Success(t *testing.T) {
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	usecase := people.NewPeopleUpdateStatusUseCase(peopleRepository)
	handler := (&PeopleHandlers{}).PatchStatus(usecase)

	newPeople := models.People{
		Status: models.StatusIn,
	}

	respRecorder := sendUpdateStatusPeopleRequest(persistance_mock.TEST_ID1, newPeople, handler)

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

	var people map[string]any
	err = json.Unmarshal(jsonbody, &people)
	if err != nil {
		t.Fatalf("invalid response unmarshal")
	}

	if people["status"].(string) != string(newPeople.Status) {
		t.Errorf("status field must be %s, got %s", string(newPeople.Status), people["status"].(string))
	}
}

func TestUpdateStatusPeopleHandler_ValidationErrors(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		people     *models.People
		statusCode int
	}{
		{
			name:       "invalid object id",
			id:         "invalid",
			people:     &models.People{},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "people id not found",
			id:         "5f8d9f1e2d862c0008e7b2f7",
			people:     &models.People{},
			statusCode: http.StatusNotFound,
		},
		{
			name:       "invalid status value",
			id:         persistance_mock.TEST_ID1,
			people:     &models.People{Status: "invalid status"},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "status is the same to current status",
			id:   persistance_mock.TEST_ID1,
			people: &models.People{
				Status: models.StatusOut,
			},
			statusCode: http.StatusConflict,
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			peopleRepository := persistance_mock.NewPeopleRepositoryMock()
			peopleRepository.Reset()

			usecase := people.NewPeopleUpdateStatusUseCase(peopleRepository)
			handler := (&PeopleHandlers{}).PatchStatus(usecase)

			respRecorder := sendUpdateStatusPeopleRequest(testItem.id, *testItem.people, handler)

			if respRecorder.Code != testItem.statusCode {
				t.Fatalf("expected status %d, got %d", testItem.statusCode, respRecorder.Code)
			}

		})
	}
}
