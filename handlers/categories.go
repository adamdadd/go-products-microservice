package handlers

import (
	"github.com/adamdadd/go-products-microservice/models"
	"log"
	"net/http"
)

type Categories struct {
	logger *log.Logger
}

func NewCategories(logger *log.Logger) *Categories {
	return &Categories{logger: logger}
}

func (c *Categories) AddCategory(writer http.ResponseWriter, request *http.Request) {
	cat := &models.Category{}

	err := cat.FromJSON(request.Body)
	if err != nil {
		http.Error(writer, "Incorrect URI", http.StatusBadRequest)
	}

	models.AddCategory(cat)
	c.logger.Printf("Product added: %#v", cat)
}
