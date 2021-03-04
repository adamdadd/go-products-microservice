package repository

import (
	"encoding/json"
	"go-products-microservice/products-api/models"
	"io"
)

type ProductRepo struct {}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

func (pr *ProductRepo) AddProduct(p *models.Product) error {
	p.ID = nextProductID()
	productList = append(productList, p)
	return nil
}

func (pr *ProductRepo) UpdateProduct(id int, p *models.Product) error {
	_, i, err := findProduct(id)
	if err != nil {
		return err
	}
	productList[i] = p
	return nil
}

func (pr *ProductRepo) DeleteProduct(id int) error {
	_, i, err := findProduct(id)
	if err != nil {
		return err
	}
	productList = append(productList[:i], productList[i+1:]...)
	return nil
}

type Products []*models.Product


func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (pr *ProductRepo) GetProducts() Products {
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