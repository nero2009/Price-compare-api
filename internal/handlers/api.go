package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"

	"github.com/gorilla/schema"
	"github.com/nero2009/pricecompare/api"
	"github.com/nero2009/pricecompare/internal/cache"
	"github.com/nero2009/pricecompare/internal/models"
)

var decoder *schema.Decoder = schema.NewDecoder()

// API is the handler for the API

func Handler(r *chi.Mux, cache *cache.Cache, db *sql.DB) {
	r.Use(chimiddle.StripSlashes)
	r.Route("/api", func(r chi.Router) {
		r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
			//get query params
			// var products api.ProductResponse

			var queryParams = api.ProductParams{}
			err := decoder.Decode(&queryParams, r.URL.Query())

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			p, err1 := GetProducts(queryParams.Product, cache, db)

			if err1 != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//add products to DB
			listingId, errm := models.CreateListing(db, queryParams.Product)

			_, err = models.CreateProduct(db, "HP Laptop", "HP Laptop 32gb 200ram", "https://www.jumia.com.ng/hp-laptop", "1000", listingId)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if errm != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(p)
		})
	})
}
