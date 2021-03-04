install:
	- go get -u github.com/gorilla/mux
	- go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	- swagger generate spec -o ./swagger.yml --scan-models

swagger-html: swagger
	- swagger serve -F=swagger swagger.yml

docker-build:
	- docker build --rm -t product-api:latest .

docker-run:
	- docker run -p 8080:8080 -d product-api:latest