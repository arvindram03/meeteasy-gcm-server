package daos

import (
	"os"
	"src/github.com/jinzhu/gorm"
	"database/sql"
	"models"
	"sync"
)

const (
	POSTGRESQL_URL = os.Getenv("HEROKU_POSTGRESQL_CHARCOAL_URL")
)

var once sync.Once
var db gorm.DB

func InitDB() {
	db, err := gorm.Open("postgres", POSTGRESQL_URL)
	if err != nil {
		return nil
	}
	dbConn := db.DB()
	dbConn.SetMaxIdleConns(2)
	dbConn.SetMaxOpenConns(10)

	db.CreateTable(&models.User{})
	db.AutoMigrate(&User{})
}

func GetDBConn() *gorm.DB{
	once.Do(InitDB)
	return &db
}
