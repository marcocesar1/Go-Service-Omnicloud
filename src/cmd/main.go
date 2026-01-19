package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/container"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoUrl := os.Getenv("MONGO_URL")
	cityApiUrl := os.Getenv("CITY_API_URL")

	app := container.NewAppContainer(mongoUrl, cityApiUrl)
	defer app.Close()

	server := http.NewServer(app)
	server.Start()
}
