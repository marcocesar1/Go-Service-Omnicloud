package repositories

import (
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
)

type PeopleRepository interface {
	Create(people *models.People) error
	FindOne(id string) (models.People, error)
	FindAll() ([]models.People, error)
	Update(id string, person *models.People) error
}
