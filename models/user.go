package models

import "time"

type User struct {
	Id             string `sql:"type:uuid" gorm:"primary_key:yes"`
	RegistrationID string `sql:"unique"`
	MobileNumber   string `sql:"unique"`
	CreatedAt      time.Time
}

func (this *User) Insert() error {
	db := GetDBConn().Create(this)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
