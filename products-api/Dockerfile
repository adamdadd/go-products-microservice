FROM golang:alpine as builder
RUN apk update && apk add build-base
WORKDIR /app
ADD . /app
RUN go build

FROM alpine
WORKDIR /app
COPY --from=builder /app/swagger.yml .
COPY --from=builder /app/products-api .
CMD ["./products-api"]
EXPOSE 8080
