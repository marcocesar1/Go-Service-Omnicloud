package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/container"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http"
)

func main() {
	_ = godotenv.Load()

	mongoUrl := os.Getenv("MONGO_URL")
	cityApiUrl := os.Getenv("CITY_API_URL")
	appPort := os.Getenv("APP_PORT")

	app := container.NewAppContainer(mongoUrl, cityApiUrl)
	defer app.Close()

	server := http.NewServer(app, appPort)
	server.Start()
}
