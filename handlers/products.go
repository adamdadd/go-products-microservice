package handlers

import (
	"context"
	"github.com/adamdadd/go-products-microservice/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := models.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "failed to process data", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(KeyProduct{}).(models.Product)

	models.AddProduct(&product)
	p.logger.Printf("Product added: %#v", product)
}

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

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, cerr := strconv.Atoi(vars["id"])
	if cerr != nil {
		p.logger.Println("Failed to convert id to int: ", vars["id"])
		http.Error(rw, "invalid URI", http.StatusBadRequest)
		return
	}

	err := models.DeleteProduct(id)
	if err != nil {
		p.logger.Println("Invalid URI for product: ", id)
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}
	p.logger.Printf("Product deleted: ", id)
}

type KeyProduct struct {}

func (p Products) ProductsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request) {
			product := models.Product{}
			err := product.FromJSON(r.Body)
			if err != nil {
				http.Error(rw, "Error trying to parse request body", http.StatusBadRequest)
				p.logger.Println("Failed parsing JSON request: ", r.Body)
			}
			ctx := context.WithValue(r.Context(), KeyProduct{}, product)
			req := r.WithContext(ctx)
			next.ServeHTTP(rw, req)
			return
	})
}