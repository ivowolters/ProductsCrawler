package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test"
	dbname   = "products"
)

type ProductDto struct {
	Url      string
	Title    string
	Keywords []string
}

func createConnection() sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return *db
}

func SaveProduct(product ProductDto) {
	db := createConnection()

	insertResult, err := db.ExecContext(
		context.Background(),
		"INSERT INTO products (company, slug, url, name) VALUES ($1, $2, $3, $4)",
		"f8765e7e-c4c6-48dd-b558-c06f0c382c6a",
		product.Title,
		product.Url,
		product.Title,
	)

	if err != nil {
		fmt.Print("Error")
		return
	}

	println(insertResult.LastInsertId())

}
