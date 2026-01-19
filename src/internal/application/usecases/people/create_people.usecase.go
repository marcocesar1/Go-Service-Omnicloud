package people

import (
	"errors"
	"fmt"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/services"
)

type CreatePeopleUseCase struct {
	peopleRepository repositories.PeopleRepository
	cityService      services.CityService
}

func NewPeopleUseCase(repo repositories.PeopleRepository, cityService services.CityService) *CreatePeopleUseCase {
	return &CreatePeopleUseCase{
		peopleRepository: repo,
		cityService:      cityService,
	}
}

func (c *CreatePeopleUseCase) Execute(people *models.People) error {
	city := c.getCity()

	people.Place = city
	people.Status = models.StatusOut

	err := c.peopleRepository.Create(people)
	if err != nil {
		if errors.Is(err, domain_err.ErrDuplicatedDoc) {
			return domain_err.DuplicatedEmail
		}

		return err
	}

	return nil
}

func (c *CreatePeopleUseCase) getCity() string {
	city, error := c.cityService.GetCityName()
	if error != nil {
		fmt.Printf("Error getting city name: %s", error.Error())

		return models.DefaultPeopleStatus
	}

	if city == "" {
		return models.DefaultPeopleStatus
	}

	return city
}
