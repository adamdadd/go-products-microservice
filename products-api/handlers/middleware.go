package handlers

import (
	"context"
	"go-products-microservice/products-api/repository"
	"net/http"
)

type KeyProduct struct {}

type KeyCategories struct {}

func (c Categories) CategoriesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut || r.Method == http.MethodPost {
			cat := repository.Category{}
			err := cat.FromJSON(r.Body)
			if err != nil {
				http.Error(rw, "Error trying to parse request body", http.StatusBadRequest)
				c.logger.Println("Failed parsing JSON request: ", r.Body)
			}
			ctx := context.WithValue(r.Context(), KeyCategories{}, cat)
			req := r.WithContext(ctx)
			next.ServeHTTP(rw, req)
			return
		}
		next.ServeHTTP(rw, r)
		return
	})
}

func (p Products) ProductsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut || r.Method == http.MethodPost {
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
		}
		next.ServeHTTP(rw, r)
		return
	})
}
