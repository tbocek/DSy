FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum* ./

# Download dependencies (if go.sum exists)
RUN go mod download || true

# Copy the source code
COPY main.go ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-server .

# Create the final image
FROM alpine:latest

# Add necessary packages
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /api-server /app/api-server

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["/app/api-server"]
