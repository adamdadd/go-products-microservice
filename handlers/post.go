package handlers

import (
	"github.com/adamdadd/go-products-microservice/models"
	"net/http"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(KeyProduct{}).(models.Product)
	models.AddProduct(&product)
	p.logger.Printf("Product added: %#v", product)
}

