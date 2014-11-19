package models

import (
	"time"
	"github.com/arvindram03/meeteasy-gcm-server/dtos"
	"github.com/arvindram03/meeteasy-gcm-server/helpers"
	"github.com/goutils/structmapper"
	"fmt"
)

type Meetup struct {
	Id             string `sql:"type:uuid"`
	MeetupTime	   time.Time `sql:"type:timestamp without timezone"`
	Location       Location
	Members        []User `gorm:"many2many:meetup_users;"`
	Description    string
}

func (this *Meetup) Insert() error{
	db := GetDBConn().Create(this)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (this *Meetup) Update() error{
	db := GetDBConn().Save(this)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (this *Meetup) FromMeetupDTO(meetupDTO *dtos.MeetupDTO) error{
	structmapper.AutoMap(*meetupDTO, this)
	this.MeetupTime = helpers.ParseTime(meetupDTO.MeetupTime)
	fmt.Println(">>>>>>>>>>>>>>>>", this.MeetupTime, meetupDTO.MeetupTime)
	db := GetDBConn().Where(meetupDTO.MemberIDs).Find(&this.Members)
	if db.Error != nil {
		return db.Error
	}
	return nil
}


type Location struct {
	Id        int64
	Latitude  float64
	Longitude float64
}
