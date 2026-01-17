package container

import (
	"log"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/usecases/people"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/repositories"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/persistance/mongo_db"
)

type AppContainer struct {
	// Infra
	Mongo *mongo_db.MongoConfig

	// Repositories
	PeopleRepository repositories.PeopleRepository

	// Use cases
	CreatePeopleUseCase *people.CreatePeopleUseCase
}

func NewAppContainer(mongoUrl string) *AppContainer {
	// Mongo
	mongo := mongo_db.NewMongoConfig(mongoUrl)
	if err := mongo.Connect(); err != nil {
		log.Fatal(err)
	}

	mongo.CreateCollections()

	// Repositories
	peopleRepo := mongo_db.NewMongoPeoplePersistence(mongo.GetDatabase())

	// Use cases
	createPeopleUseCase := people.NewPeopleUseCase(peopleRepo)

	return &AppContainer{
		Mongo:               mongo,
		PeopleRepository:    peopleRepo,
		CreatePeopleUseCase: createPeopleUseCase,
	}
}

func (a *AppContainer) Close() {
	if err := a.Mongo.Disconnect(); err != nil {
		log.Println("Error disconnecting MongoDB:", err)
	}
}
