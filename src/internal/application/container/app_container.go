package container

import (
	"log"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/city"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/persistance/mongo_db"
)

type AppContainer struct {
	// Infra
	Mongo *mongo_db.MongoConfig

	// Repositories
	PeopleRepository repositories.PeopleRepository

	// Use cases
	CreatePeopleUseCase       *people.CreatePeopleUseCase
	GetPeopleUseCase          *people.GetPeopleUseCase
	GetOnePeopleUseCase       *people.GetOnePeopleUseCase
	UpdateStatusPeopleUseCase *people.UpdateStatusPeopleUseCase
}

func NewAppContainer(mongoUrl string, cityApiUrl string) *AppContainer {
	// Mongo
	mongo := mongo_db.NewMongoConfig(mongoUrl)
	if err := mongo.Connect(); err != nil {
		log.Fatal(err)
	}

	mongo.CreateCollections()

	// Repositories
	peopleRepo := mongo_db.NewMongoPeoplePersistence(mongo.GetDatabase())

	// Services
	cityService := city.NewRandomCityApi(cityApiUrl)

	// Use cases
	createPeopleUseCase := people.NewPeopleUseCase(peopleRepo, cityService)
	getPeopleUseCase := people.NewGetPeopleUseCase(peopleRepo)
	getOnePeopleUseCase := people.NewGetOnePeopleUseCase(peopleRepo)
	updateStatusPeopleUseCase := people.NewPeopleUpdateStatusUseCase(peopleRepo)

	return &AppContainer{
		Mongo:                     mongo,
		PeopleRepository:          peopleRepo,
		CreatePeopleUseCase:       createPeopleUseCase,
		GetPeopleUseCase:          getPeopleUseCase,
		GetOnePeopleUseCase:       getOnePeopleUseCase,
		UpdateStatusPeopleUseCase: updateStatusPeopleUseCase,
	}
}

func (a *AppContainer) Close() {
	if err := a.Mongo.Disconnect(); err != nil {
		log.Println("Error disconnecting MongoDB:", err)
	}
}
