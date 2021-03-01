package handlers

import (
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
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.changeProduct(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := models.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "failed to process data", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	product := &models.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "failed to add product", http.StatusBadRequest)
	}

	p.logger.Printf("Product %#v", product)
}

func (p *Products) changeProduct(rw http.ResponseWriter, r *http.Request) {
}