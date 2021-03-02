package handlers

import (
	"github.com/adamdadd/go-products-microservice/repository"
	"net/http"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(KeyProduct{}).(repository.Product)
	repository.AddProduct(&product)
	p.logger.Printf("Product added: %#v", product)
	rw.WriteHeader(http.StatusCreated)
}

//func (c *Categories) AddCategory(writer http.ResponseWriter, request *http.Request) {
//	cat := &models.Category{}
//	err := cat.FromJSON(request.Body)
//	if err != nil {
//		http.Error(writer, "Incorrect URI", http.StatusBadRequest)
//	}
//	models.AddCategory(cat)
//	c.logger.Printf("Product added: %#v", cat)
//}
