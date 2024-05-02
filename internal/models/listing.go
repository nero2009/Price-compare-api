package models

import (
	"database/sql"
	"log"
)

// Listing is a struct that represents a listing in the database

type Listing struct {
	Id    int
	Query string
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
