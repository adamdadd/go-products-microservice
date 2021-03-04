package repository

type ProductNotFoundError struct {
	msg string
}

func (p *ProductNotFoundError) Error() string {
	return p.msg
}

func NewProductNotFoundError(msg string) error {
	return &ProductNotFoundError{msg}
}


type CategoryNotFoundError struct {
	msg string
}

func (c *CategoryNotFoundError) Error() string {
	return c.msg
}

func NewCategoryNotFoundError(msg string) error {
	return &CategoryNotFoundError{msg}
}


