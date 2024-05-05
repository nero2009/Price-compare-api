package models

import (
	"database/sql"
	"log"
)

// Listing is a struct that represents a listing in the database

type Product struct {
	Id           int
	Product_Name string
	Description  string
	Price        string
	Url          string
	Listing_Id   int64
	Created_At   string
	Updated_At   string
}

// GetListings is a function that returns all the listings in the database
func GetProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM listings")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Description, &product.Product_Name, &product.Price)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProductById(db *sql.DB, id string) (Product, error) {
	rows, err := db.Query("SELECT * FROM products WHERE id = ?", id)
	var product Product

	if err != nil {
		log.Fatal(err)
		return product, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&product.Id, &product.Description, &product.Product_Name, &product.Price)
		if err != nil {
			log.Fatal(err)
			return product, err
		}
	}
	return product, nil
}

func CreateProduct(db *sql.DB, product_name string, description string, url string, price string, listing_id int64) (int64, error) {
	result, err := db.Exec("INSERT INTO products (product_name, description, url, price, listing_id, created_at, updated_at) VALUES(?, ?, ?, ?,?,NOW(), NOW())", product_name, description, url, price, listing_id)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil

}
