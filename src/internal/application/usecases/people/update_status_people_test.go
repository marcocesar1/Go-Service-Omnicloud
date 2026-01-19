package people

import (
	"errors"
	"testing"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/persistance/persistance_mock"
)

func TestUpdateStatusPeopleUseCase_Success(t *testing.T) {
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	usecase := NewPeopleUpdateStatusUseCase(peopleRepository)

	people, err := usecase.Execute(persistance_mock.TEST_ID1, models.StatusIn)
	if err != nil {
		t.Fatalf("Error executing usecase: %s", err.Error())
	}

	if people.Status != models.StatusIn {
		t.Errorf("Status must be %s, got %s", models.StatusIn, people.Status)
	}
}

func TestUpdateStatusPeopleUseCase_ValidationErrors(t *testing.T) {
	tests := []struct {
		name          string
		id            string
		people        *models.People
		expectedError error
	}{
		{
			name:          "invalid object id",
			id:            "invalid id",
			people:        &models.People{},
			expectedError: domain_err.ErrInvalidObjectId,
		},
		{
			name: "people id not found",
			id:   "5f8d9f1e2d862c0008e7b2f7",
			people: &models.People{
				Status: models.StatusOut,
			},
			expectedError: domain_err.ErrNotFound,
		},
		{
			name: "invalid status value",
			id:   persistance_mock.TEST_ID1,
			people: &models.People{
				Status: "invalid status",
			},
			expectedError: domain_err.InvalidStatus,
		},
		{
			name: "status is the same to current status",
			id:   persistance_mock.TEST_ID1,
			people: &models.People{
				Status: models.StatusOut,
			},
			expectedError: domain_err.StatusIsTheSame,
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			repo := persistance_mock.NewPeopleRepositoryMock()

			usecase := NewPeopleUpdateStatusUseCase(repo)

			_, err := usecase.Execute(testItem.id, testItem.people.Status)

			if err == nil {
				t.Fatalf("expected error")
			}

			if !errors.Is(err, testItem.expectedError) {
				t.Fatalf("expected %v, got %v", testItem.expectedError, err)
			}
		})
	}
}
