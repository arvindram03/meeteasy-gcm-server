package models

import (
	"github.com/arvindram03/meeteasy-gcm-server/dtos"
	"github.com/goutils/structmapper"
	"time"
)

type Message struct {
	Id                int64
	UserId            string
	Latitude          float64
	Longitude         float64
	TimeToReachInMins int
	MeetupId          string
	CreatedAt         time.Time
}

func (this *Message) Insert() error {
	db := GetDBConn().Create(this)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (this *Message) FromMessageDTO(messageDTO *dtos.MessageDTO) {
	structmapper.AutoMap(*messageDTO, this)
	this.Latitude = messageDTO.Location.Latitude
	this.Longitude = messageDTO.Location.Longitude
	this.UserId = messageDTO.SenderID
}
