package handlers

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/helpers"
	"github.com/sirupsen/logrus"
)

func Setup() error {
	api := rest.NewApi()

	router, err := rest.MakeRouter(
		rest.Get("/users", getUsers),
		rest.Post("/users", postUser),
		rest.Delete("/users/:user-id", deleteUser),
	)
	if err != nil {
		helpers.Log.Error("Could not setup rest router. Err: %v", err)
		return err
	}

	api.SetApp(router)

	helpers.Log.WithFields(logrus.Fields{"port": 8080}).Info("Server is up!")

	err = http.ListenAndServe(":8088", api.MakeHandler())
	if err != nil {
		helpers.Log.Error("Could not start the server. Err: %v", err)
		return err
	}

	return nil
}
