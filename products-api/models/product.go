package models

import (
	"encoding/json"
	"io"
)

// swagger:model Product
type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	PriceGBP float32 `json:"price"`
	SKU string `json:"sku"`
	ImageURL string `json:"image_url"`
}

func NewProduct(ID int, name string, description string, priceGBP float32, SKU string, imageURL string) *Product {
	return &Product{ID: ID, Name: name, Description: description, PriceGBP: priceGBP, SKU: SKU, ImageURL: imageURL}
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Product) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

