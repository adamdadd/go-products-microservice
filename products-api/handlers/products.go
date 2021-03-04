package handlers

import (
	"go-products-microservice/products-api/repository"
	"log"
)

type Products struct {
	logger   *log.Logger
	prodRepo *repository.ProductRepo
}

func NewProducts(logger *log.Logger, pr *repository.ProductRepo) *Products {
	return &Products{logger: logger, prodRepo: pr}
}

