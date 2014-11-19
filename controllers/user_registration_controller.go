package controllers

import (
	"net/http"
	httpResponse "github.com/arvindram03/meeteasy-gcm-server/http"
	"io/ioutil"
	"encoding/json"
	"github.com/arvindram03/meeteasy-gcm-server/dtos"
	"github.com/goutils/structmapper"
	"github.com/arvindram03/meeteasy-gcm-server/models"
	"log"
)

func RegisterUser(writer http.ResponseWriter, req *http.Request) {
	byteValue, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.BadRequest(writer)
		return
	}
	userDTO := &dtos.UserDTO{}

	err = json.Unmarshal(byteValue, userDTO)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.BadRequest(writer)
		return
	}

	user := &models.User{}
	structmapper.AutoMap(*userDTO, user)

	err = user.Insert()
	if err != nil {
		httpResponse.InternalServerError(writer)
		return
	}
	httpResponse.Created(writer)
}


