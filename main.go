package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dwialim/employee-management/database"
	routes_web "github.com/dwialim/employee-management/routes"
	"github.com/joho/godotenv"

	"net/http"
)

func main() {
	db := database.InitDatabase()
	// fmt.Println("Hello World")
	server := http.NewServeMux()

	// server.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to dashboard"))
	// }))

	routes_web.MapRoutes(server, db)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("APP_NAME"))
	http.ListenAndServe(":7000", server)
}
