// Package classification Products API.
//
// Documentation for Product API.
//
//     Schemes: http
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package docs

import (
	"github.com/adamdadd/go-products-microservice/repository"
)

// swagger:route GET /products/ products listProducts
// Returns a list of products
// Responses:
//	200: productResponse

// swagger:route POST /products/ products Product
// Adds a new product
// Responses:
//	201: 201

// swagger:route PUT /products/{id} products Product
// Updates an existing product

// swagger:route DELETE /products/{id} products Product
// Deletes a product

// swagger:response
type productResponse struct {
	// in: body
	Body []repository.Product
}

// swagger:route GET /categories/ categories listCategories
// Returns a list of products
// Responses:
//	200: 201

// swagger:route POST /categories/ categories Category
// Adds a new product
// Responses:
//	201: 201

// swagger:route PUT /categories/{id} categories Category
// Updates an existing product

// swagger:route DELETE /categories/{id} categories Category
// Deletes a product

// swagger:response
type categoryResponse struct {
	// in: body
	Body []repository.Category
}

