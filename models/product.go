package models

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	PriceGBP float32 `json:"price"`
	SKU string `json:"sku"`
	ImageURL string `json:"image_url"`
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID:				1,
		Name:			"Cuppacino",
		Description: 	"A cup of the good stuff",
		PriceGBP: 		2.50,
		SKU:			"qwerty12",
		ImageURL: 		"https://google.com",
	},
	&Product{
		ID:				2,
		Name:			"Mocha",
		Description: 	"A cup of the sweet stuff",
		PriceGBP: 		2.50,
		SKU:			"qwerty13",
		ImageURL: 		"https://google.com",
	},
}