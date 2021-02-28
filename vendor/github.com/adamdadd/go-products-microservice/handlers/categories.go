package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Categories struct {
	logger *log.Logger
}

func NewCategories(logger *log.Logger) *Categories {
	return &Categories{logger: logger}
}

func (c *Categories) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "categories %s\n", r)
}
