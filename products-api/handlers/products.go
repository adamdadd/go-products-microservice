package handlers

import (
	"context"
	"github.com/adamdadd/go-products-microservice/repository"
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

type KeyProduct struct {}

func (p Products) ProductsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request) {
			product := repository.Product{}
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