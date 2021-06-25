package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/helpers"
	"github.com/sirupsen/logrus"
)

func Setup() error {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	err := http.ListenAndServe(":8088", r)
	if err != nil {
		helpers.Log.Error("Could not start the server")
		return err
	}
	helpers.Log.WithFields(logrus.Fields{"port": 8080}).Info("Server is up!")
	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: some randome text goes here just for example.\n")
}
