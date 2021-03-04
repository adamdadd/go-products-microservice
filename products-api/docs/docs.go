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
	"github.com/go-openapi/runtime/middleware"
	"go-products-microservice/products-api/models"
	"net/http"
)

// swagger:route GET /products/ products listProducts
// Returns a list of products
// Responses:
//	200: productsResponse

// swagger:route POST /products/ products Product
// Adds a new product
// Responses:
//	201: productResponse

// swagger:route PUT /products/{id} products Product
// Updates an existing product
// Responses:
//	201: productResponse

// swagger:route DELETE /products/{id} products Product
// Deletes a product
// Responses:
//	204: noBody

// swagger:response productsResponse
type productsResponse struct {
	// in: body
	Body []models.Product
}

// swagger:response productResponse
type productResponse struct {
	// in: body
	Body models.Product
}

// swagger:route GET /categories/ categories listCategories
// Returns a list of products
// Responses:
//	200: categoriesResponse

// swagger:route POST /categories/ categories Category
// Adds a new product
// Responses:
//	201: categoryResponse

// swagger:route PUT /categories/{id} categories Category
// Updates an existing product
// Responses:
//	201: categoryResponse

// swagger:route DELETE /categories/{id} categories Category
// Deletes a product
// Responses:
//	204: noBody

// swagger:response categoriesResponse
type categoriesResponse struct {
	// in: body
	Body []models.Category
}

// swagger:response categoryResponse
type categoryResponse struct {
	// in: body
	Body models.Category
}

// swagger:response noBody
type noBody struct {

}

func DocsMiddleware() http.Handler {
	opts := middleware.RedocOpts{SpecURL: "/swagger.yml"}
	sh := middleware.Redoc(opts, nil)
	return sh
}