package handlers

import (
	"github.com/adamdadd/go-products-microservice/repository"
	"net/http"
)

func (p *Products) AddProduct(writer http.ResponseWriter, request *http.Request) {
	product := request.Context().Value(KeyProduct{}).(repository.Product)
	repository.AddProduct(&product)
	writer.WriteHeader(http.StatusCreated)
	p.logger.Printf("Product added: %#v", product)
}

func (c *Categories) AddCategory(writer http.ResponseWriter, request *http.Request) {
	cat := request.Context().Value(KeyCategories{}).(repository.Category)
	repository.AddCategory(&cat)
	writer.WriteHeader(http.StatusCreated)
	c.logger.Printf("Product added: %#v", cat)
}
