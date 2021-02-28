package main

import (
	"github.com/adamdadd/go-products-microservice/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Product API", log.LstdFlags)
	hh := handlers.NewProducts(logger)
	hc := handlers.NewCategories(logger)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/", hc)

	log.Fatal(http.ListenAndServe(":8080", sm))
}
