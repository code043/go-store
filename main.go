package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

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
	//product := Product{"Book", 15.55, true}
	//pk := insertProduct(db, product)
	//fmt.Printf("ID = %d\n", pk)
	data := []Product{}
	rows, err := db.Query("SELECT name, available, price FROM product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var name string
	var price float64
	var available bool
	// query := "SELECT name, price, available FROM product WHERE id = $1"
	// err = db.QueryRow(query, 111).Scan(&name, &price, &available)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Fatalf("No rows found with ID %d", 111)
	// 	}
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Name: %s\n", name)
	// fmt.Printf("Price: %f\n", price)
	// fmt.Printf("Available: %t\n", available)
	for rows.Next() {
		err := rows.Scan(&name, &price, &available)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Product{name, price, available})
	}
	fmt.Println(data)

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
func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available) values ($1, $2, $3) RETURNING id`
	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
