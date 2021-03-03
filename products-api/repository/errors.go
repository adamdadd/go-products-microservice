package repository

import "fmt"

type RepoError interface {
	Error() string
}


type ProductNotFoundError struct {
	msg string
}
func (p ProductNotFoundError) Error() string {
	return p.msg
}


type CategoryNotFoundError struct {
	msg string
}
func (c CategoryNotFoundError) Error() string {
	return c.msg
}

var ErrorCategoryNotFound = fmt.Errorf("Category Not Found")
