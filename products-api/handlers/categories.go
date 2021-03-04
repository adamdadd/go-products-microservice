package handlers

import (
	"go-products-microservice/products-api/repository"
	"log"
)

type Categories struct {
	logger *log.Logger
	catRepo *repository.CategoryRepo
}

func NewCategories(logger *log.Logger, cr *repository.CategoryRepo) *Categories {
	return &Categories{logger: logger, catRepo: cr}
}


