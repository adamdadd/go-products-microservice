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

func nextCategoryID() int {
	lc := categoryList[len(categoryList) - 1]
	return lc.ID + 1
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

func findCategory(id int) (*Category, int, error){
	for i, c := range categoryList {
		if c.ID == id {
			return c, i, nil
		}
	}
	return nil, -1, ErrorCategoryNotFound
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
	&Category{
		ID: 			1,
		Name: 			"Tall Coffee",
		Description: 	"Double espresso coffee",
		ImageURL: 		"random.com",
	},
	&Category{
		ID: 			2,
		Name: 			"Short Coffee",
		Description: 	"Espresso coffee",
		ImageURL: 		"coffee.com",
	},
}