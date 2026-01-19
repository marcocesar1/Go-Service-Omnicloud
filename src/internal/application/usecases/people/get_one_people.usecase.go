package people

import (
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
)

type GetOnePeopleUseCase struct {
	peopleRepository repositories.PeopleRepository
}

func NewGetOnePeopleUseCase(repo repositories.PeopleRepository) *GetOnePeopleUseCase {
	return &GetOnePeopleUseCase{
		peopleRepository: repo,
	}
}

func (g *GetOnePeopleUseCase) Execute(id string) (models.People, error) {
	return g.peopleRepository.FindOne(id)
}
