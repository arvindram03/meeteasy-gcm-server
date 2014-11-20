package models

import (
	"github.com/arvindram03/meeteasy-gcm-server/dtos"
	"github.com/arvindram03/meeteasy-gcm-server/helpers"
	"github.com/goutils/structmapper"
	"time"
)

type Meetup struct {
	Id          string    `sql:"type:uuid" gorm:"primary_key:yes"`
	MeetupTime  time.Time `sql:"type:timestamp without time zone"`
	Location    Location
	LocationId  int64
	Members     []User `gorm:"many2many:meetup_users;"`
	Description string
	CreatedAt   time.Time
}

type Location struct {
	Id        int64
	Latitude  float64
	Longitude float64
}

func (this *Meetup) Insert() error {
	db := GetDBConn().Create(this)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (this *Meetup) Update() error {
	db := GetDBConn().Save(this)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (this *Meetup) FromMeetupDTO(meetupDTO *dtos.MeetupDTO) error {
	structmapper.AutoMap(*meetupDTO, this)
	this.MeetupTime = helpers.ParseTime(meetupDTO.MeetupTime)
	db := GetDBConn().Where(meetupDTO.MemberIDs).Find(&this.Members)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
