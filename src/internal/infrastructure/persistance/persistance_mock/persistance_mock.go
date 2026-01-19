package persistance_mock

import (
	"time"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const TEST_ID1 = "5f8d9f1e2d862c0008e7b2f0"
const TEST_ID2 = "5f8d9f1e2d862c0008e7b2f1"

type PeopleRepositoryMock struct {
	Error  error
	People []models.People
}

func NewPeopleRepositoryMock() *PeopleRepositoryMock {
	id1, _ := bson.ObjectIDFromHex(TEST_ID1)
	id2, _ := bson.ObjectIDFromHex(TEST_ID2)

	return &PeopleRepositoryMock{
		People: []models.People{
			{
				ID:        id1,
				Name:      "John Doe",
				Email:     "johndoe@example.com",
				Place:     "New York",
				Status:    models.StatusOut,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        id2,
				Name:      "Jane Doe",
				Email:     "janedoe@example.com",
				Place:     "New York",
				Status:    models.StatusOut,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}}
}

func (pm *PeopleRepositoryMock) Create(people *models.People) error {
	if pm.Error != nil {
		return pm.Error
	}

	// validate duplicated email
	for _, p := range pm.People {
		if p.Email == people.Email {
			return domain_err.ErrDuplicatedDoc
		}
	}

	people.ID = bson.NewObjectID()
	people.CreatedAt = time.Now()
	people.UpdatedAt = time.Now()

	pm.People = append(pm.People, *people)

	return nil
}

func (pm *PeopleRepositoryMock) FindOne(id string) (models.People, error) {
	if pm.Error != nil {
		return models.People{}, pm.Error
	}

	_, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return models.People{}, domain_err.ErrInvalidObjectId
	}

	for _, p := range pm.People {
		if p.ID.Hex() == id {
			return p, nil
		}
	}

	return models.People{}, domain_err.ErrNotFound
}

func (pm *PeopleRepositoryMock) FindAll(filters *repositories.FindAllPeopleFilter) ([]models.People, error) {
	if pm.Error != nil {
		return nil, pm.Error
	}

	return pm.People, nil
}

func (pm *PeopleRepositoryMock) Update(person *models.People) error {
	if pm.Error != nil {
		return pm.Error
	}

	for i, p := range pm.People {
		if p.ID.Hex() == person.ID.Hex() {
			pm.People[i] = *person
			return nil
		}
	}

	return nil
}
