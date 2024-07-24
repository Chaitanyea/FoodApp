package routes

import (
	"GoSQLReactProject/src/controllers"
	"net/http"
)

var FoodAppRoutes = func(router http.ServeMux) {
	router.HandleFunc("PUT /createDish", controllers.CreateDish)
	router.HandleFunc("GET /getDish/", controllers.GetAllDish)
	router.HandleFunc("GET /updateDish/:id", controllers.UpdateDish)
	router.HandleFunc("DELETE /deleteDish/:id", controllers.DeleteDish)
	router.HandleFunc("GET /getDish/:id", controllers.GetDish)
}
