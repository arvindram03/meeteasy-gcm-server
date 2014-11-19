package daos

import (
	"os"
	"github.com/jinzhu/gorm"
	"github.com/arvindram03/meeteasy-gcm-server/models"
	"sync"
	"log"
)

var once sync.Once
var db gorm.DB

func InitDB() {
	postgreSQLURL := os.Getenv("HEROKU_POSTGRESQL_CHARCOAL_URL")
	db, err := gorm.Open("postgres", postgreSQLURL)
	if err != nil {
		log.Fatal("Initialize DB: ", err)
	}
	dbConn := db.DB()
	dbConn.SetMaxIdleConns(2)
	dbConn.SetMaxOpenConns(10)

	db.CreateTable(&models.User{})
	db.AutoMigrate(&models.User{})
}

func GetDBConn() *gorm.DB{
	once.Do(InitDB)
	return &db
}
