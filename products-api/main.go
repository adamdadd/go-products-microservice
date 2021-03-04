package main

import (
	"context"
	"github.com/gorilla/mux"
	"go-products-microservice/products-api/docs"
	"go-products-microservice/products-api/handlers"
	"go-products-microservice/products-api/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "Product API", log.LstdFlags)

	productRepo := repository.NewProductRepo()
	categoryRepo := repository.NewCategoryRepo()

	productHandler := handlers.NewProducts(logger, productRepo)
	categoryHandler := handlers.NewCategories(logger, categoryRepo)

	muxRouter := mux.NewRouter()

	rootRouter := muxRouter.PathPrefix("/").Subrouter()
	rootRouter.Handle("/docs", docs.DocsMiddleware())
	rootRouter.Handle("/swagger.yml", http.FileServer(http.Dir("./")))

	productRouter := muxRouter.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("", productHandler.GetAll).Methods(http.MethodGet)
	productRouter.HandleFunc("", productHandler.Add).Methods(http.MethodPost)
	productRouter.HandleFunc("/{id:[0-9]+}", productHandler.Update).Methods(http.MethodPut)
	productRouter.HandleFunc("/{id:[0-9]+}", productHandler.Delete).Methods(http.MethodDelete)
	productRouter.Use(productHandler.ProductsMiddleware)

	categoryRouter := muxRouter.PathPrefix("/categories").Subrouter()
	categoryRouter.HandleFunc("", categoryHandler.GetAll).Methods(http.MethodGet)
	categoryRouter.HandleFunc("", categoryHandler.Add).Methods(http.MethodPost)
	categoryRouter.HandleFunc("/{id:[0-9]+}", categoryHandler.Update).Methods(http.MethodPut)
	categoryRouter.HandleFunc("/{id:[0-9]+}", categoryHandler.Delete).Methods(http.MethodDelete)
	categoryRouter.Use(categoryHandler.CategoriesMiddleware)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      muxRouter,
		IdleTimeout:  120*time.Second,
		ReadTimeout:  5*time.Second,
		WriteTimeout: 10*time.Second,
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
