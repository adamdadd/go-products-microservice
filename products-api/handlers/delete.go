package handlers

import (
	"github.com/adamdadd/go-products-microservice/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, cerr := strconv.Atoi(vars["id"])
	if cerr != nil {
		p.logger.Println("Failed to convert id to int: ", vars["id"])
		http.Error(rw, "invalid URI", http.StatusBadRequest)
		return
	}
	err := repository.DeleteProduct(id)
	if err != nil {
		p.logger.Println("Invalid URI for product: ", id)
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}
	p.logger.Printf("Product deleted: ", id)
}

