package main

import (
	"context"
	"github.com/gorilla/mux"
	"go-products-microservice/products-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "Product API", log.LstdFlags)

	productHandler := handlers.NewProducts(logger)
	categoryHandler := handlers.NewCategories(logger)

	muxRouter := mux.NewRouter()

	productRouter := muxRouter.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("/", productHandler.Get).Methods(http.MethodGet)
	productRouter.HandleFunc("/", productHandler.Add).Methods(http.MethodPost)
	productRouter.HandleFunc("/{id:[0-9]+}", productHandler.Update).Methods(http.MethodPut)
	productRouter.HandleFunc("/{id:[0-9]+}", productHandler.Delete).Methods(http.MethodDelete)
	productRouter.Use(productHandler.ProductsMiddleware)
	productRouter.UseEncodedPath()

	categoryRouter := muxRouter.PathPrefix("/categories").Subrouter()
	categoryRouter.HandleFunc("/", categoryHandler.Get).Methods(http.MethodGet)
	categoryRouter.HandleFunc("/", categoryHandler.Add).Methods(http.MethodPost)
	categoryRouter.HandleFunc("/{id:[0-9]+}", categoryHandler.Update).Methods(http.MethodPut)
	categoryRouter.HandleFunc("/{id:[0-9]+}", categoryHandler.Delete).Methods(http.MethodDelete)
	categoryRouter.Use(categoryHandler.CategoriesMiddleware)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      muxRouter,
		IdleTimeout:  120*time.Second,
		ReadTimeout:  2*time.Second,
		WriteTimeout: 2*time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, os.Kill)

	sig := <-channel

	logger.Println("Server Shutting Down: Signal:", sig)

	timeoutContext, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	err := server.Shutdown(timeoutContext)
	if err != nil {
		defer cancelFunc()
	}
}
