package models

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var DB *gorm.DB

func InitDB() {
	postgreSQLURL := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", postgreSQLURL)
	if err != nil {
		log.Fatal("Initialize DB: ", err)
	}
	DB = &db
	dbConn := db.DB()
	dbConn.SetMaxIdleConns(2)
	dbConn.SetMaxOpenConns(10)
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Meetup{}, &Location{})
}

func GetDBConn() *gorm.DB {
	if DB == nil {
		InitDB()
	}
	return DB
}
