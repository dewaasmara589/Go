package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// variable global
var DB *sql.DB

func InitDb() {
	//postgres: adalah username
	//postgres? adalah nama database

	connStr := "postgres://postgres:dewa123@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// inisialisasi untuk saign ulang
	DB = db
}
