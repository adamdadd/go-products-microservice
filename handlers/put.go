package handlers

import (
	"github.com/adamdadd/go-products-microservice/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, cerr := strconv.Atoi(vars["id"])
	if cerr != nil {
		p.logger.Println("Failed to convert id to int: ", vars["id"])
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
	}
	product := r.Context().Value(KeyProduct{}).(models.Product)
	err := models.UpdateProduct(id, &product)
	if err == models.ErrorProductNotFound || err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	p.logger.Printf("Product updated: ", id)
}

