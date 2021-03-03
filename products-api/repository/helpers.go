package repository

import "io"

type Model interface {
	FromJSON(reader io.Reader) error
	ToJSON(writer io.Writer) error
}

func nextProductID() int {
	lastElement := productList[len(productList) - 1]
	return lastElement.ID + 1
}

func nextCategoryID() int {
	lastElement := categoryList[len(categoryList) - 1]
	return lastElement.ID + 1
}

func findProduct(id int) (*Product, int, RepoError){
	for i, prod := range productList {
		if prod.ID == id {
			return prod, i, nil
		}
	}
	return nil, -1, ProductNotFoundError{"Product Not Found"}
}

func findCategory(id int) (*Category, int, RepoError){
	for i, c := range categoryList {
		if c.ID == id {
			return c, i, nil
		}
	}
	return nil, -1, CategoryNotFoundError{"Category Not Found"}
}

