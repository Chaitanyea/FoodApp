package main

import (
	"GoSQLReactProject/src/routes"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	routes.FoodAppRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
