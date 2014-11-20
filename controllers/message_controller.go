package controllers

import (
	"encoding/json"
	"github.com/arvindram03/meeteasy-gcm-server/dtos"
	httpResponse "github.com/arvindram03/meeteasy-gcm-server/http"
	"github.com/arvindram03/meeteasy-gcm-server/models"
	"io/ioutil"
	"log"
	"net/http"
)

func SendMessage(writer http.ResponseWriter, req *http.Request) {
	value, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.BadRequest(writer)
		return
	}
	messageDTO := &dtos.MessageDTO{}

	err = json.Unmarshal(value, messageDTO)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.BadRequest(writer)
		return
	}

	message := &models.Message{}
	message.FromMessageDTO(messageDTO)

	err = message.Insert()
	if err != nil {
		log.Println("[InternalServerError] ", err)
		httpResponse.BadRequest(writer)
		return
	}

}
