package main

import (
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/infrastructure/http"
)

func main() {
	server := http.NewServer()
	server.Start()
}
