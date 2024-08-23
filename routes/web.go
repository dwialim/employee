package routes_web

import (
	"database/sql"
	"net/http"

	c_employee "github.com/dwialim/employee-management/controllers"
)

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(res, req)
	})
}
func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(res, req)
	})
}
func Put(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(res, req)
	})
}
func Patch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPatch {
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(res, req)
	})
}
func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodDelete {
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(res, req)
	})
}

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	// server.HandleFunc("/", new_hello.NewIndexHello())
	// server.HandleFunc("/employee", new_employee.NewIndexEmployee())
	server.HandleFunc("/employee", c_employee.NewIndexEmployee(db))
	server.Handle("/employee/", EmployeeMux(db))
}

func EmployeeMux(db *sql.DB) http.Handler {
	EmployeeMux := http.NewServeMux()
	EmployeeMux.HandleFunc("/", c_employee.NewIndexEmployee(db))
	EmployeeMux.HandleFunc("/form", c_employee.NewFormEmployee(db))
	EmployeeMux.Handle("/delete", Get(http.HandlerFunc(c_employee.NewDeleteEmployee(db))))
	// c_employee.NewCreateEmployee(db)

	EmployeeMux.Handle("/dashboard", Get(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Dashboard v2"))
	})))
	// EmployeeMux.Handle("/profile", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	// 	res.Write([]byte("Profile"))
	// }))
	return http.StripPrefix("/employee", EmployeeMux)
}
