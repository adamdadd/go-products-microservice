FROM golang:1.16

WORKDIR /app

ADD . /app

RUN go build

CMD ["./go-products-microservice"]