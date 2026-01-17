package main

import (
	"log"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/persistance/mongo_db"
)

func main() {
	mongoConfig := mongo_db.NewMongoConfig("mongodb://root:rootpassword@localhost:27018")

	if err := mongoConfig.Connect(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := mongoConfig.Disconnect(); err != nil {
			log.Println("Error disconnecting MongoDB:", err)
		}
	}()

	mongoConfig.CreateCollections()

	// Test crud
	repo := mongo_db.NewMongoPeoplePersistence(mongoConfig.GetDatabase())

	people := models.People{
		Name:   "Marco",
		Email:  "marco@gmail.com",
		Place:  "Barcelona",
		Status: models.StatusIn,
	}

	/* err := repo.Create(&people)
	if err != nil {
		log.Fatal(err)
	} */

	people, err := repo.FindOne("696ae85a85120f6ba0a40dbf")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(people)

	people.Place = "Madrid"
	people.Status = "StatusOut"
	err = repo.Update(&people)
	if err != nil {
		log.Fatal(err)
	}

	persons, err := repo.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, person := range persons {
		log.Println(person)
	}

	//

	server := http.NewServer()
	server.Start()
}
