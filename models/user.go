package models

type User struct {
	Id             string `sql:"type:uuid"`
	RegistrationID string `sql:"unique"`
	MobileNumber   string `sql:"unique"`
}

func (this *User) Insert() error{
	db := GetDBConn().Create(this)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
