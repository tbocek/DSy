FROM golang:alpine AS builder
WORKDIR /build
COPY ../simple.go .
RUN go build simple.go

FROM alpine
WORKDIR /app
COPY --from=builder /build/simple .
ENTRYPOINT ["./simple", "8080"]