package handlers

import (
	"net/http"
)


func (p *Products) GetAll(rw http.ResponseWriter, r *http.Request) {
	listProd := p.prodRepo.GetProducts()
	err := listProd.ToJSON(rw)
	if err != nil {
		http.Error(rw, "failed to process data", http.StatusInternalServerError)
	}
}

func (c *Categories) GetAll(rw http.ResponseWriter, r *http.Request) {
	listCat := c.catRepo.GetCategories()
	err := listCat.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Failed to get categories", http.StatusInternalServerError)
	}
}

