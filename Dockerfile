FROM golang:1.24-alpine as builder
WORKDIR /app
COPY simple.go .
RUN go build -o service simple.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/service .
ENTRYPOINT ["./service", "8080"]