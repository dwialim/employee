package database

import (
	"database/sql"
	"fmt"
	"os"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDatabase() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connection := map[string]string{
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
		"user":     os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASSWORD"),
		"database": os.Getenv("DB_DATABASE"),
	}

	for name, value := range connection {
		if value == "" && name != "password" {
			log.Fatalf("Koneksi database belum di setup: %s kosong", name)
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connection["user"], connection["password"], connection["host"], connection["port"], connection["database"])
	// dsn := "remote:remote@tcp(localhost:3306)/go_app"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
