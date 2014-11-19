package services

import (
	"github.com/arvindram03/meeteasy-gcm-server/models"
	"github.com/arvindram03/meeteasy-gcm-server/daos"
)

func RegisterUser(user *models.User) error{
	return daos.RegisterUser(user)
}
