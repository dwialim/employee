package c_employee

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Employee struct {
	Id      string
	Name    string
	NPWP    string
	Address string
}

func NewIndexEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		/**
		* perbedaan Query dan Exec
		 */
		rows, err := db.Query("SELECT id, name, npwp, address FROM employee")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var employees []Employee
		for rows.Next() {
			var employee Employee

			err = rows.Scan(
				&employee.Id,
				&employee.Name,
				&employee.NPWP,
				&employee.Address,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			employees = append(employees, employee)
		}

		fp := filepath.Join("views/employee", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["employees"] = employees

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func NewFormEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			r.ParseForm()
			name := r.Form["name"][0]
			address := r.Form["address"][0]
			npwp := r.Form["npwp"][0]
			if id != "" {
				_, err := db.Exec("UPDATE employee SET name=?, npwp=?, address=? WHERE id=?", name, npwp, address, id)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			} else {
				_, err := db.Exec("INSERT INTO employee (name,npwp,address) VALUES(?,?,?)", name, npwp, address)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
			http.Redirect(w, r, "/employee/", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")

			data := make(map[string]any)
			if id != "" {
				row := db.QueryRow("SELECT name, npwp, address FROM employee WHERE id = ?", id)
				if row.Err() != nil {
					w.Write([]byte(row.Err().Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				var employee Employee
				err := row.Scan(
					&employee.Name,
					&employee.NPWP,
					&employee.Address,
				)
				employee.Id = id
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				data["employee"] = employee
			}

			fp := filepath.Join("views/employee", "form.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

func NewDeleteEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			w.Write([]byte("Data tidak ditemukan"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err := db.Exec("DELETE FROM employee WHERE id=?", id)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/employee/", http.StatusMovedPermanently)
	}
}
