package people

import (
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
)

type CreatePeopleUseCase struct {
	peopleRepository repositories.PeopleRepository
}

func NewPeopleUseCase(repo repositories.PeopleRepository) *CreatePeopleUseCase {
	return &CreatePeopleUseCase{
		peopleRepository: repo,
	}
}

func (c *CreatePeopleUseCase) Execute(people *models.People) error {
	return c.peopleRepository.Create(people)
}
