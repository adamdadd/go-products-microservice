package models

import (
	"encoding/json"
	"fmt"
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

var ErrorProductNotFound = fmt.Errorf("Product Not Found")

func AddProduct(p *Product) {
	p.ID = nextID()
	productList = append(productList, p)
}

func nextID() int {
	lp := productList[len(productList) - 1]
	return lp.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	_, i, err := findProduct(id)
	if err != nil {
		return err
	}
	productList[i] = p
	return nil
}

func DeleteProduct(id int) error {
	_, i, err := findProduct(id)
	if err != nil {
		return err
	}
	productList = append(productList[:i], productList[i+1:]...)
	return nil
}

func findProduct(id int) (*Product, int, error){
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Product) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:				1,
		Name:			"Cappuccino",
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