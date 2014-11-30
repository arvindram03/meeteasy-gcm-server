package controllers

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"github.com/arvindram03/meeteasy-gcm-server/dtos"
	httpResponse "github.com/arvindram03/meeteasy-gcm-server/http"
	"github.com/arvindram03/meeteasy-gcm-server/models"
	"github.com/goutils/structmapper"
	"io/ioutil"
	"log"
	"net/http"
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
	user.Id = uuid.New()
	err = user.Insert()
	if err != nil {
		httpResponse.InternalServerError(writer)
		return
	}
	httpResponse.Created(writer)
}

func GetRegisteredContacts(writer http.ResponseWriter, req *http.Request) {
	byteValue, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.BadRequest(writer)
		return
	}
	contactsDTO := &dtos.ContactsDTO{}

	err = json.Unmarshal(byteValue, contactsDTO)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.BadRequest(writer)
		return
	}

	contactsDTO.MobileNumbers = models.GetUsersByMobileNumbers(contactsDTO.MobileNumbers)

	response, err := json.Marshal(contactsDTO)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.InternalServerError(writer)
		return
	}

	httpResponse.StatusOKWithResponse(writer, response)
}
