package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Global variable DB
var db *gorm.DB

// Connect -> to MySQL server
func Connect() {
	dsn := "root:Google*3702@tcp(localhost:3306)/gosqlprojectschema?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Something went wrong with the DB connection : ")
		panic(err)
	}
	if d == nil {
		panic("Database could not connect, please try again!")
	}
	log.Printf("Connected to database!")
	db = d
}

// GetDb -> Getter function for DB
func GetDb() *gorm.DB {
	if db == nil {
		Connect()
	}
	return db
}
