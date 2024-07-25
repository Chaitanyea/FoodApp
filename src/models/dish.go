package models

import "gorm.io/gorm"
import "GoSQLReactProject/src/util"

var db *gorm.DB

type Dish struct {
	gorm.Model
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Rating int    `json:"rating"`
}

func initDBWithModel() {
	db = util.GetDb()
	if err := db.AutoMigrate(&Dish{}); err != nil {
		panic(err)
	}
}
