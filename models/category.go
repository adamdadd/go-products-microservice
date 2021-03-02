package models

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
	c.ID = categoryList[len(categoryList)].ID + 1
	categoryList = append(categoryList, c)
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
		1,
		"Tall Coffee",
		"Double espresso coffee",
		"random.com",
	},
	&Category{
		2,
		"Short Coffee",
		"Finest espresso",
		"coffeee.com",
	},
}