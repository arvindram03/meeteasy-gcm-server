package controllers

import (
	"net/http"
	httpResponse "github.com/arvindram03/meeteasy-gcm-server/http"
	"io/ioutil"
	"encoding/json"
	"github.com/arvindram03/meeteasy-gcm-server/dtos"
	"github.com/goutils/structmapper"
	"github.com/arvindram03/meeteasy-gcm-server/models"
	"github.com/arvindram03/meeteasy-gcm-server/services"
)

func Register(writer http.ResponseWriter, req *http.Request) {
	value, err := ioutil.ReadAll(req.Body)
	if err != nil {
		httpResponse.BadRequest(writer)
		return
	}
	var userDTO *dtos.UserDTO

	err = json.Unmarshal(value, userDTO)
	if err != nil {
		httpResponse.BadRequest(writer)
		return
	}

	var user *models.User
	structmapper.AutoMap(userDTO, user)

	err = services.RegisterUser(user)
	if err != nil {
		httpResponse.InternalServerError(writer)
		return
	}
	httpResponse.Created(writer)
	return
}
