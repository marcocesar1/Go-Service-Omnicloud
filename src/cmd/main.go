package main

import (
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/application/container"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http"
)

func main() {
	mongoUrl := "mongodb://root:rootpassword@localhost:27018"

	app := container.NewAppContainer(mongoUrl)
	defer app.Close()

	server := http.NewServer(app)
	server.Start()
}
