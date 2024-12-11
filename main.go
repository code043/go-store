package main

import (
	"database/sql"
	"log"
)

func main() {
	conn := "postgres://postgres:postgres@localhost:5432/go-store?sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
