package handlers

import (
	"go-products-microservice/products-api/repository"
	"net/http"
)


func (p *Products) Get(rw http.ResponseWriter, r *http.Request) {
	lp := repository.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "failed to process data", http.StatusInternalServerError)
	}
}

func (c *Categories) Get(writer http.ResponseWriter, request *http.Request) {
	lc := repository.GetCategories()
	err := lc.ToJSON(writer)
	if err != nil {
		http.Error(writer, "Failed to get categories", http.StatusInternalServerError)
	}
}

