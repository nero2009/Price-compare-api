package handlers

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/nero2009/pricecompare/api"
	log "github.com/sirupsen/logrus"
)

func GetJumiaProducts() []api.Product {
	url := "https://www.jumia.com.ng/catalog/?q=laptop"

	response, err := http.Get(url)

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

func GetProductsFromKonga() []api.Product {
	url := "https://kara.com.ng/catalogsearch/result/?q=laptop"

	response, err := http.Get(url)

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
