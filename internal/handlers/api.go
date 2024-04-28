package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"

	"github.com/gorilla/schema"
	"github.com/nero2009/pricecompare/api"
	"github.com/nero2009/pricecompare/internal/cache"
	log "github.com/sirupsen/logrus"
)

var decoder *schema.Decoder = schema.NewDecoder()

// API is the handler for the API

func Handler(r *chi.Mux, cache *cache.Cache) {
	r.Use(chimiddle.StripSlashes)
	r.Route("/api", func(r chi.Router) {
		r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
			//get query params
			var products api.ProductResponse

			var queryParams = api.ProductParams{}
			err := decoder.Decode(&queryParams, r.URL.Query())

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Printf("%v", queryParams)

			cachedProducts, isPresent := cache.Get("products")

			if isPresent {
				json.Unmarshal(cachedProducts.([]byte), &products)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(products)
				return
			}

			kongaProducts := GetProductsFromKonga()
			jumiaProducts := GetJumiaProducts()

			products = api.ProductResponse{
				JumiaProducts: jumiaProducts,
				KongaProducts: kongaProducts,
			}

			jsonProducts, err := json.Marshal(products)

			if err != nil {
				log.Error("Error marshalling the products", err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//cache the products in redis

			cache.Set("products", jsonProducts, 10*time.Minute)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(products)

		})
	})
}
