package people

import (
	"errors"
	"testing"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/persistance/persistance_mock"
)

func TestGetOnePeopleUseCase_Success(t *testing.T) {
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	usecase := NewGetOnePeopleUseCase(peopleRepository)

	id := persistance_mock.TEST_ID1
	people, err := usecase.Execute(id)
	if err != nil {
		t.Fatalf("Error executing usecase: %s", err.Error())
	}

	if people.ID.Hex() != id {
		t.Errorf("ID field must be %s, got %s", id, people.ID.Hex())
	}
}

func TestGetOnePeopleUseCase_ValidationErrors(t *testing.T) {
	tests := []struct {
		name          string
		id            string
		expectedError error
	}{
		{
			name:          "invalid object id",
			id:            "invalid",
			expectedError: domain_err.ErrInvalidObjectId,
		},
		{
			name:          "people not found",
			id:            "5f8d9f1e2d862c0008e7b2f3",
			expectedError: domain_err.ErrNotFound,
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			repo := persistance_mock.NewPeopleRepositoryMock()

			usecase := NewGetOnePeopleUseCase(repo)

			_, err := usecase.Execute(testItem.id)

			if err == nil {
				t.Fatalf("expected error")
			}

			if !errors.Is(err, testItem.expectedError) {
				t.Fatalf("expected %v, got %v", testItem.expectedError, err)
			}
		})
	}
}
