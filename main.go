package main

import (
	"log"
	"net/http"
	"os"
	"github.com/adamdadd/go-products-microservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "Product API", log.LstdFlags)
	hh := handlers.NewProduct(l)
	http.HandleFunc("/", hh)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
