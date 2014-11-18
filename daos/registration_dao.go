package daos

import "models"

func RegisterUser(user *models.User) error{
	db := GetDBConn().Create(user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

