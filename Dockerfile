# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o hawklens cmd/hawklens/main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Install necessary runtime libraries
RUN apk add --no-cache ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/hawklens .
COPY --from=builder /app/dashboard ./dashboard

# Expose API port
EXPOSE 8080

# Environment variables
ENV GIN_MODE=release

# Run the API server by default
CMD ["./hawklens", "serve"]
