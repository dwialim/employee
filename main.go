package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dwialim/employee/database"
	"github.com/joho/godotenv"

	routes_web "github.com/dwialim/employee/routes"
	"net/http"
)

func main() {
	database.InitDatabase()
	// fmt.Println("Hello World")
	server := http.NewServeMux()

	routes_web.MapRoutes(server)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("APP_NAME"))
	http.ListenAndServe(":7000", server)
}
