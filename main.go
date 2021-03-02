package main

import (
	"context"
	_ "github.com/adamdadd/go-products-microservice/docs"
	"github.com/adamdadd/go-products-microservice/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "Product API", log.LstdFlags)
	ph := handlers.NewProducts(logger)
	ch := handlers.NewCategories(logger)

	sm := mux.NewRouter()

	gr := sm.Methods("GET").Subrouter()
	gr.HandleFunc("/products/", ph.GetProducts)
	gr.HandleFunc("/categories/", ch.GetCategories)

	ppr := sm.Methods("POST").Subrouter()
	ppr.HandleFunc("/products/", ph.AddProduct)
	ppr.Use(ph.ProductsMiddleware)

	putr := sm.Methods("PUT").Subrouter()
	putr.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	putr.Use(ph.ProductsMiddleware)

	dr := sm.Methods("DELETE").Subrouter()
	dr.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	s := &http.Server{
		Addr: ":8080",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, os.Kill)

	sig := <-channel

	logger.Println("shutting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
