package handlers

import (
	"encoding/json"
	"github.com/adamdadd/go-products-microservice/models"
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	lp := models.GetProducts()
	data, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "failed to process data", http.StatusInternalServerError)
	}

}
