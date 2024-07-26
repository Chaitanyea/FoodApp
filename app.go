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
	handler := enableCORS(router)
	http.Handle("/", handler)
	util.Connect()
	models.Init()
	routes.FoodAppRoutes(router)
	// Enable CORS
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// enableCORS enables Cross-Origin Resource Sharing (CORS) for the given handler.
func enableCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "content-Type, content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		log.Printf("CORS is enabled!")
		w.WriteHeader(http.StatusOK)
		handler.ServeHTTP(w, r)
	})
}
