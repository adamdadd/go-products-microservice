package handlers

import (
	"log"
)

type Categories struct {
	logger *log.Logger
}


func NewCategories(logger *log.Logger) *Categories {
	return &Categories{logger: logger}
}

