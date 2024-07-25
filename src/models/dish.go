package models

import (
	"GoSQLReactProject/src/util"
	"fmt"
	"gorm.io/gorm"
)

var db *gorm.DB

type Dish struct {
	gorm.Model
	Name   string  `json:"name"`
	Price  int     `json:"price"`
	Rating float32 `json:"rating"`
}

// Init populate global Db for class, initialize with model
func Init() {
	db = util.GetDb()
	if err := db.AutoMigrate(&Dish{}); err != nil {
		panic(err)
	}
}

// CreateDish in Database
func (d *Dish) CreateDish() *Dish {
	if db == nil {
		fmt.Println("Db is nil")
	}
	db.Create(&d)
	return d
}

// GetAllDishes from Database
func GetAllDishes() []Dish {
	var dishes []Dish
	db.Find(&dishes)
	return dishes
}

// GetDish from Database
func GetDish(Key int64) *Dish {
	var getDish Dish
	db.First(&getDish, Key)
	return &getDish
}

// DeleteDish from Database
func DeleteDish(Key int64) *Dish {
	dish := GetDish(Key)
	db.Delete(&Dish{}, Key)
	return dish
}
