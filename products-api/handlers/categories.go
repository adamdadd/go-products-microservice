package handlers

import (
	"context"
	"github.com/adamdadd/go-products-microservice/repository"
	"log"
	"net/http"
)

type Categories struct {
	logger *log.Logger
}

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
func NewCategories(logger *log.Logger) *Categories {
	return &Categories{logger: logger}
}

