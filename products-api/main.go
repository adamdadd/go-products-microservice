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
	pr.HandleFunc("/", ph.Get).Methods(http.MethodGet)
	pr.HandleFunc("/", ph.Add).Methods(http.MethodPost)
	pr.HandleFunc("/{id:[0-9]+}", ph.Update).Methods(http.MethodPut)
	pr.HandleFunc("/{id:[0-9]+}", ph.Delete).Methods(http.MethodDelete)
	pr.Use(ph.ProductsMiddleware)

	cr := sm.PathPrefix("/categories").Subrouter()
	cr.HandleFunc("/", ch.Get).Methods(http.MethodGet)
	cr.HandleFunc("/", ch.Add).Methods(http.MethodPost)
	cr.HandleFunc("/{id:[0-9]+}", ch.Update).Methods(http.MethodPut)
	cr.HandleFunc("/{id:[0-9]+}", ch.Delete).Methods(http.MethodDelete)
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

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	err := s.Shutdown(tc)
	if err != nil {
		cancel()
	}
}
