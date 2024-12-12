package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	conn := "postgres://postgres:dev@localhost:5432/go_store?sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	createProductTable(db)
}
func createProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product(
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6, 2) NOT NULL,
		available BOOLEAN,
		created timestamp DEFAULT NOW()
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
