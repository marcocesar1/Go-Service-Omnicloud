package people

import (
	"testing"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/persistance/persistance_mock"
)

func TestGetPeopleUseCase_Success(t *testing.T) {
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	usecase := NewGetPeopleUseCase(peopleRepository)

	people, err := usecase.Execute(&repositories.FindAllPeopleFilter{})
	if err != nil {
		t.Fatalf("Error executing usecase: %s", err.Error())
	}

	if len(people) == 0 {
		t.Errorf("People must be greater than 0")
	}
}

func TestGetPeopleUseCase_Success_Filter_Status(t *testing.T) {
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	usecase := NewGetPeopleUseCase(peopleRepository)

	people, err := usecase.Execute(&repositories.FindAllPeopleFilter{
		Status: string(models.StatusIn),
	})
	if err != nil {
		t.Fatalf("Error executing usecase: %s", err.Error())
	}

	if len(people) != 1 {
		t.Errorf("People length must be 1, got %d", len(people))
	}
}
