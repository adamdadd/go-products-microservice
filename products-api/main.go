package main

import (
	"context"
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

	pr := sm.PathPrefix("/products").Subrouter()
	pr.HandleFunc("/", ph.GetProducts).Methods(http.MethodGet)
	pr.HandleFunc("/", ph.AddProduct).Methods(http.MethodPost)
	pr.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct).Methods(http.MethodPut)
	pr.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct).Methods(http.MethodDelete)
	pr.Use(ph.ProductsMiddleware)

	cr := sm.PathPrefix("/categories").Subrouter()
	cr.HandleFunc("/", ch.GetCategories).Methods(http.MethodGet)
	cr.HandleFunc("/", ch.AddCategory).Methods(http.MethodPost)
	cr.Use(ch.CategoriesMiddleware)

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

	logger.Println("Server Shutting Down: Signal:", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
