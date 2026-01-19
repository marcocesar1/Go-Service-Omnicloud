package people

import (
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
)

type GetPeopleUseCase struct {
	peopleRepository repositories.PeopleRepository
}

func NewGetPeopleUseCase(repo repositories.PeopleRepository) *GetPeopleUseCase {
	return &GetPeopleUseCase{
		peopleRepository: repo,
	}
}

func (g *GetPeopleUseCase) Execute(filters *repositories.FindAllPeopleFilter) ([]models.People, error) {
	return g.peopleRepository.FindAll(filters)
}
