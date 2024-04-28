package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/nero2009/pricecompare/api"
	"github.com/nero2009/pricecompare/internal/cache"
	log "github.com/sirupsen/logrus"
)

func GetJumiaProducts(query string) []api.Product {
	//check if the product is in the cache
	//if it is, return the product
	baseURL, err := url.Parse("https://www.jumia.com.ng/catalog")
	if err != nil {
		log.Error("Error parsing the URL", err)
		return nil
	}
	queryParams := url.Values{}
	queryParams.Set("q", query)
	baseURL.RawQuery = queryParams.Encode()

	response, err := http.Get(baseURL.String())

	if err != nil {
		log.Error("Error fetching the URL", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Error("Error fetching the URL", response.Status)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Error("Error parsing HTML", err)
	}

	var products []api.Product

	document.Find("article.prd").Each(func(index int, element *goquery.Selection) {
		name := element.Find("div.info h3.name").Text()
		price := element.Find("div.info div.prc").Text()
		id := element.Find("a.core").AttrOr("data-gtm-id", "")
		products = append(products, api.Product{
			Id:    id,
			Name:  name,
			Price: price,
		})

		fmt.Println(id)
	})
	return products
}

func GetProductsFromKonga(query string) []api.Product {
	baseURL, err := url.Parse("https://kara.com.ng/catalogsearch/result")
	if err != nil {
		log.Error("Error parsing the URL", err)
		return nil
	}
	queryParams := url.Values{}
	queryParams.Set("q", query)
	baseURL.RawQuery = queryParams.Encode()

	response, err := http.Get(baseURL.String())
	if err != nil {
		log.Error("Error fetching the URL", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Error("Error fetching the URL", response.Status)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Error("Error parsing HTML", err)
		return nil
	}

	var products []api.Product

	document.Find("li.item.product.product-item").Each(func(index int, element *goquery.Selection) {
		name := element.Find("a.product-item-link").Text()
		price := element.Find("span.price").Text()
		id := element.Find("div.price-box.price-final_price").AttrOr("data-product-id", "")
		products = append(products, api.Product{
			Id:    id,
			Name:  name,
			Price: price,
		})

		fmt.Println(products)

	})

	return products
}

func GetProducts(query string, cache *cache.Cache) (api.ProductResponse, error) {
	var products api.ProductResponse

	cachedProducts, isPresent := cache.Get(query)

	if isPresent {
		json.Unmarshal(cachedProducts.([]byte), &products)
		return products, nil
	}

	kongaProducts := GetProductsFromKonga(query)
	jumiaProducts := GetJumiaProducts(query)

	products = api.ProductResponse{
		JumiaProducts: jumiaProducts,
		KongaProducts: kongaProducts,
	}

	jsonProducts, err := json.Marshal(products)

	if err != nil {
		log.Error("Error marshalling the products", err)
		return api.ProductResponse{}, err
	}

	//cache the products in redis
	cache.Set(query, jsonProducts, 10*time.Minute)
	return products, nil
}
