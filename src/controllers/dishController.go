package controllers

import (
	"GoSQLReactProject/src/models"
	"GoSQLReactProject/src/util"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GetAllDish from database
func GetAllDish(responseWriter http.ResponseWriter, request *http.Request) {
	Dishes := models.GetAllDishes()
	res, _ := json.Marshal(Dishes)
	_, err := responseWriter.Write(res)
	//Handle any error associated with Write
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Add("Content-Type", "application/json")
}

// GetDish by Id stored in Database
func GetDish(responseWriter http.ResponseWriter, request *http.Request) {
	vars := request.PathValue("id")
	ID, err := strconv.ParseInt(vars, 0, 0)
	//Handle error associated with Integer Parsing
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	Dish := models.GetDish(ID)
	//Dish not found in database
	if Dish.ID == 0 {
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte("DishID Not Found"))
		return
	}
	res, _ := json.Marshal(Dish)
	_, err = responseWriter.Write(res)
	//Handle any error associated with Write
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Add("Content-Type", "application/json")
}

// CreateDish create a new dish using dish model
func CreateDish(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize a new instance of Dish
	newDish := models.Dish{}
	// Decode the JSON request body into the newDish instance
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&newDish)
	//Decoding error
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	Dish := newDish.CreateDish()
	res, err := json.Marshal(Dish)
	//Marshalling error
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	responseWriter.WriteHeader(http.StatusCreated)
	_, err = responseWriter.Write(res)
	//Handle error associated with Write
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
}

// DeleteDish from DB
func DeleteDish(responseWriter http.ResponseWriter, request *http.Request) {
	vars := request.PathValue("id")
	ID, err := strconv.ParseInt(vars, 0, 0)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	Dish := models.DeleteDish(ID)
	//When Dish is not found in database
	if Dish.ID == 0 {
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte("DishID Not Found"))
		return
	}
	res, _ := json.Marshal(Dish)
	_, err = responseWriter.Write(res)
	//Handle error associated with Write
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Add("Content-Type", "application/json")
}

// UpdateDish call to DeleteDish and then create new dish
func UpdateDish(responseWriter http.ResponseWriter, request *http.Request) {
	var updateDish = models.Dish{}
	err := util.ParseJson(request, &updateDish)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	vars := request.PathValue("id")
	ID, err := strconv.ParseInt(vars, 0, 0)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	//Delete Dish from DB
	Dish := models.DeleteDish(ID)
	//Dish not found in database
	if Dish.ID == 0 {
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte("DishID Not Found, please create a new dish"))
		return
	}
	//Create New dish
	updateDish.CreateDish()
	res, _ := json.Marshal(updateDish)
	_, err = responseWriter.Write(res)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	responseWriter.WriteHeader(http.StatusAccepted)
	responseWriter.Header().Add("Content-Type", "application/json")
}
