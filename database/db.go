package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	var err error

	// Inisialisasi koneksi ke database
	DB, err = sql.Open("mysql", "root:@/godatabase")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Tes koneksi
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection established")
}

func GetDB() *sql.DB {
	return DB
}