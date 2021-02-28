package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("hello world")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "something went wrong", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "%s\n", d)
}
