# Start from the official Golang image
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies (if any)
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o server main.go

# Use a minimal image for the final container
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Expose port 8080
EXPOSE 8080

# Run the server
CMD ["./server"]
