package people

import (
	"errors"
	"testing"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/city"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/persistance/persistance_mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestCreatePeopleUseCase_Success(t *testing.T) {
	cityService := city.NewRandomCityApiMock()
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	usecase := NewPeopleUseCase(peopleRepository, cityService)

	people := &models.People{
		Name:  "John",
		Email: "johndoe_city@example.com",
	}

	err := usecase.Execute(people)
	if err != nil {
		t.Fatalf("Error executing usecase: %s", err.Error())
	}

	if people.Place != "New York" {
		t.Errorf("Place field must be New York, got %s", people.Place)
	}

	if (people.ID == bson.ObjectID{}) {
		t.Errorf("ID field must not be empty")
	}
}

func TestCreatePeopleUseCase_Success_With_Default_Place(t *testing.T) {
	cityService := city.NewRandomCityApiMock()
	peopleRepository := persistance_mock.NewPeopleRepositoryMock()

	cityService.Error = errors.New("error getting city name")

	usecase := NewPeopleUseCase(peopleRepository, cityService)

	people := &models.People{
		Name:  "John",
		Email: "johndoe_unknown@example.com",
	}

	err := usecase.Execute(people)
	if err != nil {
		t.Fatalf("Error executing usecase: %s", err.Error())
	}

	if people.Place != models.DefaultPeopleStatus {
		t.Errorf("Place field must be %s, got %s", models.DefaultPeopleStatus, people.Place)
	}

	if people.ID.IsZero() {
		t.Errorf("ID field must not be empty")
	}
}

func TestCreatePeopleUseCase_ValidationErrors(t *testing.T) {
	tests := []struct {
		name          string
		people        *models.People
		expectedError error
	}{
		{
			name: "empty name",
			people: &models.People{
				Name:  "",
				Email: "test@test.com",
			},
			expectedError: domain_err.ErrNameRequired,
		},
		{
			name: "short name",
			people: &models.People{
				Name:  "jo",
				Email: "test@test.com",
			},
			expectedError: domain_err.ErrNameInvalidLength,
		},
		{
			name: "empty email",
			people: &models.People{
				Name:  "John",
				Email: "",
			},
			expectedError: domain_err.ErrEmailRequired,
		},
		{
			name: "invalid email",
			people: &models.People{
				Name:  "John",
				Email: "test",
			},
			expectedError: domain_err.ErrEmailInvalid,
		},
		{
			name: "duplicated email",
			people: &models.People{
				Name:  "John",
				Email: "johndoe@example.com",
			},
			expectedError: domain_err.DuplicatedEmail,
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			cityService := city.NewRandomCityApiMock()
			repo := persistance_mock.NewPeopleRepositoryMock()

			usecase := NewPeopleUseCase(repo, cityService)

			err := usecase.Execute(testItem.people)

			if err == nil {
				t.Fatalf("expected error")
			}

			if !errors.Is(err, testItem.expectedError) {
				t.Fatalf("expected %v, got %v", testItem.expectedError, err)
			}
		})
	}
}
