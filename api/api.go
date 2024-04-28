package api

type Product struct {
	Id    string
	Name  string
	Price string
}

type ProductParams struct {
	Product string
}

type ProductResponse struct {
	JumiaProducts []Product
	KongaProducts []Product
}
