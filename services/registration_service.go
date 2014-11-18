package services

import (
	"models"
	"daos"
)

func RegisterUser(user *models.User) error{
	return daos.RegisterUser(user)
}
