package models

import (
	"database/sql"
	"log"
)

// Listing is a struct that represents a listing in the database

type Listing struct {
	Id        int
	Query     string
	CreatedAt string
	UpdatedAt string
}

// GetListings is a function that returns all the listings in the database
func GetListings(db *sql.DB) []Listing {
	rows, err := db.Query("SELECT * FROM listings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var listings []Listing
	for rows.Next() {
		var listing Listing
		err := rows.Scan(&listing.Id, &listing.Query)
		if err != nil {
			log.Fatal(err)
		}
		listings = append(listings, listing)
	}
	return listings
}

// get Listing by id
func GetListingById(db *sql.DB, id string) (Listing, error) {
	rows, err := db.Query("SELECT * FROM LISTING WHERE id = ?", id)
	var listings Listing

	if err != nil {
		log.Fatal(err)
		return listings, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&listings.Id, &listings.Query)
		if err != nil {
			log.Fatal(err)
			return listings, err
		}
	}
	return listings, nil
}

// create a new listing
func CreateListing(db *sql.DB, query string) (int64, error) {
	result, err := db.Exec("INSERT INTO listing (query, created_at, updated_at) VALUES(?, NOW(), NOW())", query)
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
