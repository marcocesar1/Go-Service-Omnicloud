package people

import (
	"errors"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/validations"
)

type UpdateStatusPeopleUseCase struct {
	peopleRepository repositories.PeopleRepository
}

func NewPeopleUpdateStatusUseCase(repo repositories.PeopleRepository) *UpdateStatusPeopleUseCase {
	return &UpdateStatusPeopleUseCase{
		peopleRepository: repo,
	}
}
func (u *UpdateStatusPeopleUseCase) Execute(id string, status models.PeopleStatus) (*models.People, error) {
	people, err := u.peopleRepository.FindOne(id)
	if err != nil {
		if errors.Is(err, domain_err.ErrNotFound) {
			return nil, domain_err.ErrNotFound
		}
		return nil, err
	}

	err = validations.ValidateStatus(status)
	if err != nil {
		return nil, err
	}

	if people.Status == status {
		return nil, domain_err.StatusIsTheSame
	}

	people.Status = status
	err = u.peopleRepository.Update(&people)
	if err != nil {
		return nil, err
	}

	return &people, nil
}
