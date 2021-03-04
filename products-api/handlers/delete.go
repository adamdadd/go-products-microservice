package handlers

import "C"
import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)


func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, cerr := strconv.Atoi(vars["id"])
	if cerr != nil {
		p.logger.Println("Failed to convert id to int: ", vars["id"])
		http.Error(rw, "invalid URI", http.StatusBadRequest)
		return
	}

	err := p.prodRepo.DeleteProduct(id)
	if err != nil {
		p.logger.Println("Invalid URI for product: ", id, "\nError:", err)
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	p.logger.Printf("Product deleted: ", id)
	rw.WriteHeader(http.StatusNoContent)
}


func (c *Categories) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, cerr := strconv.Atoi(vars["id"])
	if cerr != nil {
		c.logger.Println("Failed to convert id to int: ", vars["id"])
		http.Error(rw, "invalid URI", http.StatusBadRequest)
		return
	}

	err := c.catRepo.DeleteCategory(id)
	if err != nil {
		c.logger.Println("Invalid URI for category: ", id, "\nError:", err)
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	c.logger.Printf("Category deleted: ", id)
	rw.WriteHeader(http.StatusNoContent)
}

