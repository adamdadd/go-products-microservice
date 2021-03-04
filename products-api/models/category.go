package models

import (
	"encoding/json"
	"io"
)

// swagger:model Category
type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
}

func NewCategory(ID int, name string, description string, imageURL string) *Category {
	return &Category{ID: ID, Name: name, Description: description, ImageURL: imageURL}
}


func (c *Category) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(c)
}

func (c *Category) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(c)
}

