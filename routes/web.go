package routes_web

import (
	new_employee "go-app-crud/controllers"
	// new_hello "go-app-crud/controllers/hello-world"
	"net/http"
)

func MapRoutes(server *http.ServeMux) {
	// server.HandleFunc("/", new_hello.NewIndexHello())
	server.HandleFunc("/employee", new_employee.NewIndexEmployee())
	server.HandleFunc("/dwialim", new_employee.NewIndexEmployee())
	// server.HandleFunc("/testing", new_employee.NewIndexEmployee())
}