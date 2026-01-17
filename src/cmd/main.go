package main

import (
	"log"

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

	server := http.NewServer()
	server.Start()
}
