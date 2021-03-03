package repository

import (
	"encoding/json"
	"fmt"
	"io"
)

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
}

var ErrorCategoryNotFound = fmt.Errorf("Category Not Found")

func (c *Category) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(c)
}

func (c *Category) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(c)
}

func AddCategory(c *Category) {
	c.ID = nextCategoryID()
	categoryList = append(categoryList, c)
}

func UpdateCategory(id int, c *Category) error {
	_, i, err := findCategory(id)
	if err != nil {
		return err
	}
	categoryList[i] = c
	return nil
}

func DeleteCategory(id int) error {
	_, i, err := findCategory(id)
	if err != nil {
		return err
	}
	categoryList = append(categoryList[:i], categoryList[i+1:]...)
	return nil
}

type Categories []*Category

func (c *Categories) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(c)
}

func GetCategories() Categories {
	return categoryList
}

var categoryList = []*Category{
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