package controllers

import (
	"GoSQLReactProject/src/models"
	"GoSQLReactProject/src/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllDish(responseWriter http.ResponseWriter, request *http.Request) {
	Dishes := models.GetAllDishes()
	res, _ := json.Marshal(Dishes)
	_, err := responseWriter.Write(res)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
}

func GetDish(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	dishId := vars["id"]
	ID, err := strconv.Atoi(dishId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	Dish := models.GetDish(int64(ID))
	res, _ := json.Marshal(Dish)
	_, err = responseWriter.Write(res)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
}

func CreateDish(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize a new instance of Dish
	newDish := models.Dish{}

	// Decode the JSON request body into the newDish instance
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&newDish)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	Dish := newDish.CreateDish()
	res, err := json.Marshal(Dish)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	responseWriter.Write(res)
}

func DeleteDish(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	dishId := vars["id"]
	ID, err := strconv.Atoi(dishId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	Dish := models.DeleteDish(int64(ID))
	res, _ := json.Marshal(Dish)
	_, err = responseWriter.Write(res)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
}

func UpdateDish(responseWriter http.ResponseWriter, request *http.Request) {
	var updateDish = models.Dish{}
	err := util.ParseJson(request, &updateDish)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	vars := mux.Vars(request)
	dishId := vars["id"]
	ID, err := strconv.Atoi(dishId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	models.DeleteDish(int64(ID))
	updateDish.CreateDish()
	res, _ := json.Marshal(updateDish)
	_, err = responseWriter.Write(res)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusAccepted)
}
