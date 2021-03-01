module github.com/adamdadd/go-products-microservice/handlers

go 1.16

require (
	github.com/adamdadd/go-products-microservice/models v0.0.0-20210301111305-ba0373ba5281
	github.com/gorilla/mux v1.8.0
)

replace github.com/adamdadd/go-products-microservice/models v0.0.0-20210301111305-ba0373ba5281 => ../models
