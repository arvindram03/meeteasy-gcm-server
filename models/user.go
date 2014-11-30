package models

import (
	"time"
	"fmt"
)

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

func GetUsersByMobileNumbers(mobileNumbers []string) []string {
	db := GetDBConn()
	var registeredMobileNumbers []string
	value := db.Model(&User{}).Where("mobile_number in (?)", mobileNumbers).Pluck("mobile_number", &registeredMobileNumbers)
	fmt.Println(">>>>>>>>>>", value)
	return registeredMobileNumbers
}
