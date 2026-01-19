package repositories

import (
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
)

type FindAllPeopleFilter struct {
	Status string
}

type PeopleRepository interface {
	Create(people *models.People) error
	FindOne(id string) (models.People, error)
	FindAll(filters *FindAllPeopleFilter) ([]models.People, error)
	Update(person *models.People) error
}
