package controllers

import (
	httpResponse "github.com/arvindram03/meeteasy-gcm-server/http"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/goutils/structmapper"
	"github.com/arvindram03/meeteasy-gcm-server/dtos"
	"github.com/arvindram03/meeteasy-gcm-server/models"
	"code.google.com/p/go-uuid/uuid"
)

func CreateMeetup(writer http.ResponseWriter, req *http.Request) {
	value, err := ioutil.ReadAll(req.Body)
	if err != nil {
		httpResponse.BadRequest(writer)
		return
	}
	meetupDTO := &dtos.MeetupDTO{}

	err = json.Unmarshal(value, meetupDTO)
	if err != nil {
		httpResponse.BadRequest(writer)
		return
	}

	meetup := &models.Meetup{}
	structmapper.AutoMap(meetupDTO, meetup)

	meetup.Id = uuid.New()
	err = meetup.Insert()
	if err != nil {
		httpResponse.InternalServerError(writer)
		return
	}

	httpResponse.Created(writer)
}
