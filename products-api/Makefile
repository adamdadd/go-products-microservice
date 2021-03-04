install:
	- go get -u github.com/gorilla/mux
	- go get -u github.com/go-swagger/go-swagger/cmd/swagger
	- go get -u github.com/go-openapi/runtime/middleware

swagger:
	- swagger generate spec -o ./swagger.yml --scan-models

swagger-html: swagger
	- swagger serve -F=swagger swagger.yml

docker-build:
	- docker build --rm -t product-api:latest .

docker-run: docker-build
	- docker run -p 8080:8080 -d product-api:latest