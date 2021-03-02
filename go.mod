module github.com/adamdadd/go-products-microservice

go 1.16

require (
	github.com/adamdadd/go-products-microservice/handlers v0.0.0-20210301112624-178f47fc47cb
	github.com/adamdadd/go-products-microservice/models v0.0.0-20210301111305-ba0373ba5281 // indirect
	github.com/go-openapi/errors v0.20.0 // indirect
	github.com/go-openapi/validate v0.20.2 // indirect
	github.com/go-swagger/go-swagger v0.26.1 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/spf13/afero v1.5.1 // indirect
	golang.org/x/mod v0.4.1 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210301091718-77cc2087c03b // indirect
	golang.org/x/tools v0.1.0 // indirect
)

replace github.com/adamdadd/go-products-microservice/handlers v0.0.0-20210301112624-178f47fc47cb => ./handlers

replace github.com/adamdadd/go-products-microservice/models v0.0.0-20210301111305-ba0373ba5281 => ./models
