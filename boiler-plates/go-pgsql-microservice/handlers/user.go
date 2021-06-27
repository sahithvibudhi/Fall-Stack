package handlers

import (
	"fmt"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/helpers"
	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/models"
)

func getUsers(w rest.ResponseWriter, r *rest.Request) {
	var users []models.User
	models.DB.Find(&users)
	w.WriteJson(users)
}

func postUser(w rest.ResponseWriter, r *rest.Request) {
	user := models.User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		helpers.Log.Error("Error in creating a user. Err: %v", err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}

	models.DB.Insert(&user)

	w.WriteHeader(http.StatusCreated)
	w.WriteJson(user)
}

func deleteUser(w rest.ResponseWriter, r *rest.Request) {
	userId := r.PathParam("user-id")
	user := models.User{}

	affected, err := models.DB.ID(userId).Delete(&user)
	if err != nil {
		helpers.Log.Error(fmt.Scanf("Could not delete user. Err: %s", err.Error()))
	}

	if affected == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
