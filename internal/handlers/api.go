package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"

	"github.com/gorilla/schema"
	"github.com/nero2009/pricecompare/api"
	"github.com/nero2009/pricecompare/internal/cache"
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

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(p)
		})
	})
}
