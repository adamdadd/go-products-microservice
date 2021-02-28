package handlers

import (
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
	rw.Write([]byte("categories"))
}
