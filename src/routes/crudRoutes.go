package routes

import (
	"GoSQLReactProject/src/controllers"
	"github.com/gorilla/mux"
)

var FoodAppRoutes = func(router *mux.Router) {
	router.HandleFunc("/createDish", controllers.CreateDish).Methods("POST")
	router.HandleFunc("/getDish/", controllers.GetAllDish).Methods("GET")
	router.HandleFunc("/updateDish/:id", controllers.UpdateDish).Methods("PUT")
	router.HandleFunc("/deleteDish/:id", controllers.DeleteDish).Methods("DELETE")
	router.HandleFunc("/getDish/:id", controllers.GetDish).Methods("GET")
}
