package main

import (
	"os"

	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/handlers"
	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/helpers"
	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/models"
)

func main() {
	// Load configuration from .env file
	// Copy from .example.env and make changes
	err := helpers.LoadEnv()
	if err != nil {
		os.Exit(1)
	}

	// Connection to DB and run migrations
	err = models.Setup()
	if err != nil {
		os.Exit(1)
	}

	// setup backend routes and serve HTTP server
	handlers.Setup()
}
