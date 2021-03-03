package handlers

import (
	"github.com/adamdadd/go-products-microservice/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, cerr := strconv.Atoi(vars["id"])
	if cerr != nil {
		p.logger.Println("Failed to convert id to int: ", vars["id"])
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
	}
	product := r.Context().Value(KeyProduct{}).(repository.Product)
	err := repository.UpdateProduct(id, &product)
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	p.logger.Printf("Product updated: ", id)
}

func (c *Categories) Update(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, cerr := strconv.Atoi(vars["id"])
	if cerr != nil {
		c.logger.Println("Failed to convert id to int: ", vars["id"])
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
	}
	cat := r.Context().Value(KeyCategories{}).(repository.Category)
	err := repository.UpdateCategory(id, &cat)
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	c.logger.Printf("Product updated: ", id)
}
