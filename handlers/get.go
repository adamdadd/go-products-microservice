package handlers

import (
	"github.com/adamdadd/go-products-microservice/models"
	"net/http"
)


func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := models.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "failed to process data", http.StatusInternalServerError)
	}
}

func (c *Categories) GetCategories(writer http.ResponseWriter, request *http.Request) {
	lc := models.GetCategories()
	err := lc.ToJSON(writer)
	if err != nil {
		http.Error(writer, "Failed to get categories", http.StatusInternalServerError)
	}
}

