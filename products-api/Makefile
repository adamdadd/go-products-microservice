install:
	- go get -u github.com/gorilla/mux
	- GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	- swagger generate spec -o ./swagger.yml --scan-models

swagger-html: swagger
	- swagger serve -F=swagger swagger.yml