package database

import (
	"log"

	_ "github.com/mattn/go-sqlite3"

	"database/sql"
)

// NewDB init the database pool
func NewDB(dbFile string) *sql.DB {

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	return db
}
