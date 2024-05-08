package handlers

import (
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

func Handler(r *chi.Mux, cache *cache.Cache) {
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
			p, err := GetProducts(queryParams.Product, cache)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//add products to DB
			listingId, err := models.CreateListing(queryParams.Product)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//loop through the products and add them to the database

			for _, product := range p.JumiaProducts {
				_, err = models.CreateProduct(product.Name, product.Price, "https://www.jumia.com.ng/catalog", "1000", listingId)

				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}

			for _, product := range p.KongaProducts {
				_, err = models.CreateProduct(product.Name, product.Price, "https://kara.com.ng/catalogsearch/result", "1000", listingId)

				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(p)
		})
	})
}
