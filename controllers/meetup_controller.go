package controllers

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"github.com/arvindram03/meeteasy-gcm-server/dtos"
	httpResponse "github.com/arvindram03/meeteasy-gcm-server/http"
	"github.com/arvindram03/meeteasy-gcm-server/models"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateMeetup(writer http.ResponseWriter, req *http.Request) {
	value, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.BadRequest(writer)
		return
	}
	meetupDTO := &dtos.MeetupDTO{}

	err = json.Unmarshal(value, meetupDTO)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.BadRequest(writer)
		return
	}

	meetup := &models.Meetup{}
	err = meetup.FromMeetupDTO(meetupDTO)
	if err != nil {
		httpResponse.InternalServerError(writer)
		return
	}
	meetup.Id = uuid.New()
	err = meetup.Insert()
	if err != nil {
		log.Println("[InternalServerError] ", err)
		httpResponse.InternalServerError(writer)
		return
	}

	httpResponse.Created(writer)
}

func UpdateMeetup(writer http.ResponseWriter, req *http.Request) {
	value, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.BadRequest(writer)
		return
	}
	meetupDTO := &dtos.MeetupDTO{}

	err = json.Unmarshal(value, meetupDTO)
	if err != nil {
		log.Println("[ParseError] ", err)
		httpResponse.BadRequest(writer)
		return
	}

	meetup := &models.Meetup{}
	err = meetup.FromMeetupDTO(meetupDTO)
	if err != nil {
		httpResponse.InternalServerError(writer)
		return
	}
	params := req.URL.Query()
	meetup.Id = params.Get(":meetupId")
	err = meetup.Update()
	if err != nil {
		httpResponse.InternalServerError(writer)
		return
	}
	httpResponse.StatusOK(writer)
}
