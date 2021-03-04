package models

import "io"

type Model interface {
	FromJSON(reader io.Reader) error
	ToJSON(writer io.Writer) error
}

