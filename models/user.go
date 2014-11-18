package models

type User struct {
	RegistrationID string
	MobileNumber string	`gorm:"primary_key:yes"`
}

