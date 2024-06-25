package routes_web

import (
	new_employee "github.com/dwialim/employee/controllers"
	// new_hello "github.com/dwialim/employee/controllers/hello-world"
	"net/http"
)

func MapRoutes(server *http.ServeMux) {
	// server.HandleFunc("/", new_hello.NewIndexHello())
	server.HandleFunc("/employee", new_employee.NewIndexEmployee())
	server.HandleFunc("/dwialim", new_employee.NewIndexEmployee())
	// server.HandleFunc("/testing", new_employee.NewIndexEmployee())
}
