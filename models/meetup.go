package models

type Meetup struct {
	Id             string `gorm:"type:uuid"`
	Location       Location
	MobileNumber   string
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


type Location struct {
	Id        int64
	Latitude  float64
	Longitude float64
}
