package controllers

import (
	"net/http"
	httpResponse "http"
	"io/ioutil"
	"encoding/json"
	"dtos"
	"github.com/goutils/structmapper"
	"models"
	"services"
)

func Register(writer http.ResponseWriter, req *http.Request) {
	value, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return httpResponse.BadRequest(writer)
	}
	var userDTO *dtos.UserDTO

	err = json.Unmarshal(value, userDTO)
	if err != nil {
		return httpResponse.BadRequest(writer)
	}

	var user *models.User
	structmapper.AutoMap(userDTO, user)

	err = services.RegisterUser(user)
	if err != nil {
		return httpResponse.InternalServerError(writer)
	}

	return httpResponse.Created(writer)
}
