# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go files
COPY go.mod ./
RUN go mod download

COPY . .

# Build the binary
RUN go build -o myapp main.go

# Stage 2: Minimal runtime image
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary and HTML template from builder
COPY --from=builder /app/myapp .
COPY --from=builder /app/index.html .

# Expose port
EXPOSE 8080

# Run the app
CMD ["./myapp"]
