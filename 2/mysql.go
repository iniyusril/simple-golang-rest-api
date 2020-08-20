package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golang")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
