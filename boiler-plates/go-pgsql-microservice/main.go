package main

import (
	"os"

	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/handlers"
	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/models"
)

func main() {
	// @TODO load the configuration
	err := models.Setup()
	if err != nil {
		os.Exit(1)
	}
	handlers.Setup()
}
