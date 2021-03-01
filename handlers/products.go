package handlers

import (
	"github.com/adamdadd/go-products-microservice/models"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
		re := regexp.MustCompile("/([0-9]+)")
		g := re.FindAllStringSubmatch(r.URL.Path, -1)
		idStr := g[0][1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			p.logger.Println("Unable tot convert id to int", idStr)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProduct(id, rw, r)
		p.logger.Println("Updated product", id)
	}

	if r.Method == http.MethodDelete {
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

	models.AddProduct(product)
	p.logger.Printf("Product added: %#v", product)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	product := &models.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "failed to add product", http.StatusBadRequest)
	}

	err = models.UpdateProduct(id, product)
	if err == models.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}

func (p *Products) deleteProduct(id int, rw http.ResponseWriter, r *http.Request) {
	
}