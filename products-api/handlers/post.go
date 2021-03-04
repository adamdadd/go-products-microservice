package handlers

import (
	"go-products-microservice/products-api/models"
	"net/http"
)


func (p *Products) Add(rw http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(KeyProduct{}).(models.Product)

	repoErr := p.prodRepo.AddProduct(&product)
	if repoErr != nil {
		http.Error(rw, "failed to add product", http.StatusInternalServerError)
		p.logger.Printf("Post failed to add product: %#v", product)
	}

	rw.WriteHeader(http.StatusCreated)
	responseErr := product.ToJSON(rw)
	if responseErr != nil {
		http.Error(rw, "Updated product response failed", http.StatusInternalServerError)
		p.logger.Printf("Post response failed: %#v", product)
	}

	p.logger.Printf("Product added: %#v", product)
}


func (c *Categories) Add(rw http.ResponseWriter, r *http.Request) {
	cat := r.Context().Value(KeyCategories{}).(models.Category)

	repoErr := c.catRepo.AddCategory(&cat)
	if repoErr != nil {
		http.Error(rw, "failed to add category", http.StatusInternalServerError)
		c.logger.Printf("Post failed to add category: %#v", cat)
	}

	rw.WriteHeader(http.StatusCreated)
	responseErr := cat.ToJSON(rw)
	if responseErr != nil {
		http.Error(rw, "Updated category response failed", http.StatusInternalServerError)
		c.logger.Printf("Post response failed: %#v", cat)
	}

	c.logger.Printf("Category added: %#v", cat)
}
