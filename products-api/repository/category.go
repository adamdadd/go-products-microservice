package repository

import (
	"encoding/json"
	"io"
)

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
}

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