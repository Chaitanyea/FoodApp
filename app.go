package main

import (
	"GoSQLReactProject/src/models"
	"GoSQLReactProject/src/routes"
	"GoSQLReactProject/src/util"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	util.Connect()
	models.Init()
	routes.FoodAppRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
