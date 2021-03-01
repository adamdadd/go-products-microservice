module github.com/adamdadd/go-products-microservice

go 1.16


require (
	github.com/adamdadd/go-products-microservice/handlers v0.0.0-20210301112624-178f47fc47cb
	github.com/adamdadd/go-products-microservice/models v0.0.0-20210301111305-ba0373ba5281 // indirect
)
replace github.com/adamdadd/go-products-microservice/handlers v0.0.0-20210301112624-178f47fc47cb => ./handlers
replace github.com/adamdadd/go-products-microservice/models v0.0.0-20210301111305-ba0373ba5281 => ./models
