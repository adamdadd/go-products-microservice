FROM golang:alpine as builder
WORKDIR /app
ADD . /app
RUN go build

FROM alpine
WORKDIR /app
COPY --from=builder /app .
CMD ["./email-service"]
EXPOSE 8080