FROM golang:alpine as builder
WORKDIR /app
ADD . /app
RUN go build

FROM alpine
WORKDIR /app
COPY --from=builder /app .
CMD ["./products-api"]
EXPOSE 8080