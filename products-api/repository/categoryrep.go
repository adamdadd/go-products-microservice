package repository

import (
	"encoding/json"
	"go-products-microservice/products-api/models"
	"io"
)

type CategoryRepo struct {}

func NewCategoryRepo() *CategoryRepo {
	return &CategoryRepo{}
}

func (cr *CategoryRepo) AddCategory(c *models.Category) error {
	c.ID = nextCategoryID()
	categoryList = append(categoryList, c)
	return nil
}

func (cr *CategoryRepo) UpdateCategory(id int, c *models.Category) error {
	_, i, err := findCategory(id)
	if err != nil {
		return err
	}
	categoryList[i] = c
	return nil
}

func (cr *CategoryRepo) DeleteCategory(id int) error {
	_, i, err := findCategory(id)
	if err != nil {
		return err
	}
	categoryList = append(categoryList[:i], categoryList[i+1:]...)
	return nil
}

func (cr *CategoryRepo) GetCategories() Categories {
	return categoryList
}

type Categories []*models.Category

func (c *Categories) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(c)
}

var categoryList = Categories {
	{
		ID: 			1,
		Name: 			"Tall Coffee",
		Description: 	"Double espresso coffee",
		ImageURL: 		"random.com",
	},
	{
		ID: 			2,
		Name: 			"Short Coffee",
		Description: 	"Espresso coffee",
		ImageURL: 		"coffee.com",
	},
}