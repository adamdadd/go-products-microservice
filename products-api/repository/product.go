package repository

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	PriceGBP float32 `json:"price"`
	SKU string `json:"sku"`
	ImageURL string `json:"image_url"`
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Product) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func AddProduct(p *Product) {
	p.ID = nextProductID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, i, err := findProduct(id)
	if err != nil {
		return err
	}
	productList[i] = p
	return nil
}

func DeleteProduct(id int) RepoError {
	_, i, err := findProduct(id)
	if err != nil {
		return err
	}
	productList = append(productList[:i], productList[i+1:]...)
	return nil
}

type Products []*Product


func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = Products{
	{
		ID:				1,
		Name:			"Cappuccino",
		Description: 	"A cup of the good stuff",
		PriceGBP: 		2.50,
		SKU:			"qwerty12",
		ImageURL: 		"https://google.com",
	},
	{
		ID:				2,
		Name:			"Mocha",
		Description: 	"A cup of the sweet stuff",
		PriceGBP: 		2.50,
		SKU:			"qwerty13",
		ImageURL: 		"https://google.com",
	},
}