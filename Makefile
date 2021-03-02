install:
	- GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	- GO111MODULE=off swagger generate spec -o ./swagger.yml --scan-models

swagger-html: swagger
	- swagger serve -F=swagger swagger.yml