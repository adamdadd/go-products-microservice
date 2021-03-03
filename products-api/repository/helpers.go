package repository

import "io"

type Model interface {
	FromJSON(reader io.Reader) error
	ToJSON(writer io.Writer) error
}

func nextProductID() int {
	lp := productList[len(productList) - 1]
	return lp.ID + 1
}

func nextCategoryID() int {
	lc := categoryList[len(categoryList) - 1]
	return lc.ID + 1
}

func findProduct(id int) (*Product, int, error){
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}

func findCategory(id int) (*Category, int, error){
	for i, c := range categoryList {
		if c.ID == id {
			return c, i, nil
		}
	}
	return nil, -1, ErrorCategoryNotFound
}

