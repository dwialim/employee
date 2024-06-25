package main

import (
	"go-app-crud/database"
	routes_web "go-app-crud/routes"
	"net/http"
)

func main() {
	database.InitDatabase()
	// fmt.Println("Hello World")
	server := http.NewServeMux()

	routes_web.MapRoutes(server)

	http.ListenAndServe(":7000", server)
}
