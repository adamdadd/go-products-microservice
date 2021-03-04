package repository

import (
	"go-products-microservice/products-api/models"
)

type CategoryRepository interface {
	AddCategory(c *models.Category) error
	UpdateCategory(id int, c *models.Category) error
	DeleteCategory(id int) error
	GetCategories() Categories
}

type ProductRepository interface {
	AddProduct(c *models.Product) error
	UpdateProduct(id int, c *models.Product) error
	DeleteProduct(id int) error
	GetProducts() Products
}

func nextProductID() int {
	lastElement := productList[len(productList) - 1]
	return lastElement.ID + 1
}

func nextCategoryID() int {
	lastElement := categoryList[len(categoryList) - 1]
	return lastElement.ID + 1
}

func findProduct(id int) (*models.Product, int, error){
	for i, prod := range productList {
		if prod.ID == id {
			return prod, i, nil
		}
	}
	return nil, -1, NewProductNotFoundError("Product not found")
}

func findCategory(id int) (*models.Category, int, error){
	for i, c := range categoryList {
		if c.ID == id {
			return c, i, nil
		}
	}
	return nil, -1, NewCategoryNotFoundError("Category not found")
}

