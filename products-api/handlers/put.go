package handlers

import (
	"github.com/gorilla/mux"
	"go-products-microservice/products-api/models"
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

	product := r.Context().Value(KeyProduct{}).(models.Product)

	err := p.prodRepo.UpdateProduct(id, &product)
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	p.logger.Printf("Product updated: ", id)
	rw.WriteHeader(http.StatusCreated)
	responseError := product.ToJSON(rw)
	if responseError != nil {
		http.Error(rw, "Failed to response with updated product", http.StatusInternalServerError)
		return
	}
}


func (c *Categories) Update(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, cerr := strconv.Atoi(vars["id"])
	if cerr != nil {
		c.logger.Println("Failed to convert id to int: ", vars["id"])
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
	}

	cat := r.Context().Value(KeyCategories{}).(models.Category)

	err := c.catRepo.UpdateCategory(id, &cat)
	if err != nil {
		http.Error(rw, "Category not found", http.StatusNotFound)
		return
	}

	c.logger.Printf("Category updated: ", id)

	rw.WriteHeader(http.StatusCreated)
	responseError := cat.ToJSON(rw)
	if responseError != nil {
		http.Error(rw, "Failed to response with updated category", http.StatusInternalServerError)
		return
	}
}