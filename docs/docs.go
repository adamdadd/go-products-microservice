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
	"github.com/adamdadd/go-products-microservice/models"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// Responses:
//	200: productResponse
type productResponse struct {
	// in: body
	Body []models.Product
}

